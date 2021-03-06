// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package ricapie2

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	e2tapi "github.com/onosproject/onos-api/go/onos/e2t/e2"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-pci/pkg/southbound/admin"
	"github.com/onosproject/onos-pci/pkg/store"
	app "github.com/onosproject/onos-ric-sdk-go/pkg/config/app/default"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"
	configutils "github.com/onosproject/onos-ric-sdk-go/pkg/config/utils"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	sdkSub "github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
	"sync"
	"time"
)

const RCPreServiceModelOIDV1 = "1.3.6.1.4.1.53148.1.1.2.100"

var log = logging.GetLogger("sb-ricapie2")

const (
	ServiceModelName       = "oran-e2sm-rc-pre"
	ServiceModelVersion    = "v1"
	ReportPeriodConfigPath = "/report_period/interval"
)

// E2Session is responsible for mapping connections to and interactions with the northbound of ONOS-E2T
type E2Session struct {
	E2SubEndpoint  string
	E2SubInstance  sdkSub.Context
	SubDelTrigger  chan bool
	E2TEndpoint    string
	RicActionID    types.RicActionID
	ReportPeriodMs uint64
	AppConfig      *app.Config
	mu             sync.RWMutex
	configEventCh  chan event.Event
}

// NewSession creates a new southbound session of ONOS-KPIMON
func NewSession(e2tEndpoint string, e2subEndpoint string, ricActionID int32, reportPeriodMs uint64) *E2Session {
	log.Info("Creating RicAPIE2Session")
	return &E2Session{
		E2SubEndpoint:  e2subEndpoint,
		E2TEndpoint:    e2tEndpoint,
		RicActionID:    types.RicActionID(ricActionID),
		ReportPeriodMs: reportPeriodMs,
	}
}

// Run starts the southbound to watch indication messages
func (s *E2Session) Run(indChan chan *store.E2NodeIndication, ctrlReqChans map[string]chan *e2tapi.ControlRequest, adminSession *admin.E2AdminSession) {
	log.Info("Started KPIMON Southbound session")
	s.configEventCh = make(chan event.Event)
	go func() {
		_ = s.watchConfigChanges()
	}()
	s.SubDelTrigger = make(chan bool)
	s.manageConnections(indChan, ctrlReqChans, adminSession)
}

func (s *E2Session) updateReportPeriod(event event.Event) error {
	interval, err := s.AppConfig.Get(event.Key)
	if err != nil {
		return err
	}
	value, err := configutils.ToUint64(interval.Value)
	if err != nil {
		return err
	}
	s.mu.Lock()
	s.ReportPeriodMs = value
	s.mu.Unlock()
	return nil
}

func (s *E2Session) processConfigEvents() {
	for configEvent := range s.configEventCh {
		if configEvent.Key == ReportPeriodConfigPath {
			log.Debug("Report Period: Config Event received:", configEvent)
			err := s.updateReportPeriod(configEvent)
			if err != nil {
				log.Error(err)
			}
			err = s.deleteSuscription()
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func (s *E2Session) watchConfigChanges() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := s.AppConfig.Watch(ctx, s.configEventCh)
	if err != nil {
		return err
	}
	s.processConfigEvents()
	return nil
}

func (s *E2Session) deleteSuscription() error {
	err := s.E2SubInstance.Close()
	if err != nil {
		log.Error(err)
	}
	s.SubDelTrigger <- true
	return err
}

// manageConnections handles connections between ONOS-PCI and ONOS-E2T/E2Sub.
func (s *E2Session) manageConnections(indChan chan *store.E2NodeIndication, ctrlReqChans map[string]chan *e2tapi.ControlRequest, adminSession *admin.E2AdminSession) {
	for {
		nodeIDs, err := adminSession.GetListE2NodeIDs()
		if err != nil {
			log.Errorf("Cannot get NodeIDs through Admin API: %s", err)
			continue
		} else if len(nodeIDs) == 0 {
			log.Warn("CU-CP is not running - wait until CU-CP is ready")
			time.Sleep(1000 * time.Millisecond)
			continue
		}
		log.Infof("Received E2Nodes: %v", nodeIDs)
		var wg sync.WaitGroup
		for _, id := range nodeIDs {
			hasOID, err := s.checkOID(id, adminSession)
			if err != nil {
				log.Error(err)
				continue
			} else if !hasOID {
				log.Warnf("E2Node %v does not support RC Pre Service model (OID: %v)", id, RCPreServiceModelOIDV1)
				continue
			}

			log.Infof("E2Node %v supports RC Pre service - trying to send subscription message", id)
			if _, ok := ctrlReqChans[id]; !ok {
				ctrlReqChans[id] = make(chan *e2tapi.ControlRequest)
				log.Infof("CtrlReqChans: %v", ctrlReqChans)
			}

			wg.Add(1)
			go func(id string, wg *sync.WaitGroup) {
				defer wg.Done()
				for {
					s.manageConnection(indChan, id, ctrlReqChans[id])
				}
			}(id, &wg)
		}
		wg.Wait()
		time.Sleep(1000 * time.Millisecond) // retry timer
	}
}

func (s *E2Session) checkOID(nodeID string, session *admin.E2AdminSession) (bool, error) {
	ranFunctions, err := session.GetRANFunctions(nodeID)
	if err != nil {
		return false, err
	}

	for _, ranFunction := range ranFunctions {
		log.Debugf("ranFunction.Oid (%v): %v", ranFunction.Description, ranFunction.Oid)
		if ranFunction.Oid == RCPreServiceModelOIDV1 {
			return true, nil
		}
	}
	return false, nil
}

func (s *E2Session) manageConnection(indChan chan *store.E2NodeIndication, nodeID string, ctrlReqChan chan *e2tapi.ControlRequest) {
	err := s.subscribeE2T(indChan, nodeID, ctrlReqChan)
	if err != nil {
		log.Warn("Error happens when subscription %s", err)
	}
}

func (s *E2Session) createSubscriptionRequest(nodeID string) (subscription.SubscriptionDetails, error) {

	return subscription.SubscriptionDetails{
		E2NodeID: subscription.E2NodeID(nodeID),
		ServiceModel: subscription.ServiceModel{
			Name:    ServiceModelName,
			Version: ServiceModelVersion,
		},
		EventTrigger: subscription.EventTrigger{
			Payload: subscription.Payload{
				Encoding: subscription.Encoding_ENCODING_PROTO,
				Data:     s.createEventTriggerData(),
			},
		},
		Actions: []subscription.Action{
			{
				ID:   int32(s.RicActionID),
				Type: subscription.ActionType_ACTION_TYPE_REPORT,
				SubsequentAction: &subscription.SubsequentAction{
					Type:       subscription.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
					TimeToWait: subscription.TimeToWait_TIME_TO_WAIT_ZERO,
				},
			},
		},
	}, nil
}

func (s *E2Session) createEventTriggerData() []byte {
	log.Infof("Received period value: %v", s.ReportPeriodMs)

	//e2smRcEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinitionPeriodic(int32(s.ReportPeriodMs))
	// use reactive way in this stage - for the future, we can choose one of two options: proactive or reactive
	e2smRcEventTriggerDefinition, err := pdubuilder.CreateE2SmRcPreEventTriggerDefinitionUponChange()
	if err != nil {
		log.Errorf("Failed to create event trigger definition data: %v", err)
		return []byte{}
	}

	err = e2smRcEventTriggerDefinition.Validate()
	if err != nil {
		log.Errorf("Failed to validate the event trigger definition: %v", err)
		return []byte{}
	}

	protoBytes, err := proto.Marshal(e2smRcEventTriggerDefinition)
	if err != nil {
		log.Errorf("Failed to marshal event trigger definition: %v", err)
	}

	return protoBytes
}

func (s *E2Session) subscribeE2T(indChan chan *store.E2NodeIndication, nodeID string, ctrlReqChan chan *e2tapi.ControlRequest) error {
	log.Infof("Connecting to ONOS-E2Sub...%s", s.E2SubEndpoint)

	e2SubHost := strings.Split(s.E2SubEndpoint, ":")[0]
	e2SubPort, err := strconv.Atoi(strings.Split(s.E2SubEndpoint, ":")[1])
	if err != nil {
		log.Error("onos-e2sub's port information or endpoint information is wrong.")
		return err
	}
	e2tHost := strings.Split(s.E2TEndpoint, ":")[0]
	e2tPort, err := strconv.Atoi(strings.Split(s.E2TEndpoint, ":")[1])
	if err != nil {
		log.Error("onos-e2t's port information or endpoint information is wrong.")
		return err
	}

	clientConfig := e2client.Config{
		AppID: "onos-pci",
		E2TService: e2client.ServiceConfig{
			Host: e2tHost,
			Port: e2tPort,
		},
		SubscriptionService: e2client.ServiceConfig{
			Host: e2SubHost,
			Port: e2SubPort,
		},
	}

	client, err := e2client.NewClient(clientConfig)

	if err != nil {
		log.Warn("Can't open E2Client.")
		return err
	}

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	subReq, err := s.createSubscriptionRequest(nodeID)
	if err != nil {
		log.Warn("Can't create SubsdcriptionRequest message")
		return err
	}

	s.E2SubInstance, err = client.Subscribe(ctx, subReq, ch)
	if err != nil {
		log.Warn("Can't send SubscriptionRequest message")
		return err
	}

	log.Infof("Start forwarding Indication message to PCI controller")
	for {
		select {
		case indMsg := <-ch:
			go func() {
				indChan <- &store.E2NodeIndication{
					NodeID: nodeID,
					IndMsg: indMsg,
				}
			}()
		case ctrlReqMsg := <-ctrlReqChan:
			log.Infof("Received E2Node: %v, Session E2Node: %v - Raw message: %v", ctrlReqMsg.E2NodeID, nodeID, ctrlReqMsg)
			if string(ctrlReqMsg.E2NodeID) != nodeID {
				log.Errorf("E2Node ID does not match: E2Node ID E2Session - %v; E2Node ID in Ctrl Message - %v", nodeID, ctrlReqMsg.E2NodeID)
				continue
			}
			ctrlRespMsg, err := client.Control(ctx, ctrlReqMsg)
			if err != nil {
				log.Errorf("Failed to send control message - %v", err)
			} else if ctrlRespMsg == nil {
				log.Errorf("Control response message is nil")
			}

			ctrlAck := ctrlRespMsg.GetControlAcknowledge()
			ctrlFailure := ctrlRespMsg.GetControlFailure()

			if ctrlAck != nil {
				ctrlOutcome := &e2sm_rc_pre_ies.E2SmRcPreControlOutcome{}
				err = proto.Unmarshal(ctrlAck.GetControlOutcome(), ctrlOutcome)
				if err != nil {
					log.Errorf("Failed to get control outcome - %v", err)
				}

				log.Infof("Received ACK message %v", ctrlOutcome.GetControlOutcomeFormat1())
			}
			if ctrlFailure != nil {
				log.Errorf("Control Failure message arrived")
			}

		case trigger := <-s.SubDelTrigger:
			if trigger {
				log.Info("Reset indChan to close subscription")
				return nil
			}
		}
	}

}

type E2SmRcPreControlHandler struct {
	NodeID              string
	ServiceModelName    e2tapi.ServiceModelName
	ServiceModelVersion e2tapi.ServiceModelVersion
	ControlMessage      []byte
	ControlHeader       []byte
	ControlAckRequest   e2tapi.ControlAckRequest
	EncodingType        e2tapi.EncodingType
}

func (c *E2SmRcPreControlHandler) CreateRcControlRequest() (*e2tapi.ControlRequest, error) {
	controlRequest := &e2tapi.ControlRequest{
		E2NodeID: e2tapi.E2NodeID(c.NodeID),
		Header: &e2tapi.RequestHeader{
			EncodingType: c.EncodingType,
			ServiceModel: &e2tapi.ServiceModel{
				Name:    c.ServiceModelName,
				Version: c.ServiceModelVersion,
			},
		},
		ControlAckRequest: c.ControlAckRequest,
		ControlHeader:     c.ControlHeader,
		ControlMessage:    c.ControlMessage,
	}
	return controlRequest, nil
}

func (c *E2SmRcPreControlHandler) CreateRcControlHeader(cellID uint64, cellIDLen uint32, priority int32, plmnID []byte) ([]byte, error) {
	eci := &e2sm_rc_pre_ies.BitString{
		Value: cellID,
		Len:   cellIDLen,
	}
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlHeader(priority, plmnID, eci)
	if err != nil {
		return []byte{}, err
	}

	err = newE2SmRcPrePdu.Validate()
	if err != nil {
		return []byte{}, err
	}

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	if err != nil {
		return []byte{}, err
	}

	return protoBytes, nil
}

func (c *E2SmRcPreControlHandler) CreateRcControlMessage(ranParamID int32, ranParamName string, ranParamValue int32) ([]byte, error) {
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlMessage(ranParamID, ranParamName, ranParamValue)
	if err != nil {
		return []byte{}, err
	}

	err = newE2SmRcPrePdu.Validate()
	if err != nil {
		return []byte{}, err
	}

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	if err != nil {
		return []byte{}, err
	}

	return protoBytes, nil
}
