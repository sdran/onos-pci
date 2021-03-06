// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package admin

import (
	"context"
	adminapi "github.com/onosproject/onos-api/go/onos/e2t/admin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/southbound"
	"google.golang.org/grpc"
	"io"
	"time"
)

var log = logging.GetLogger("sb-admin")

// E2AdminSession is responsible for mapping connections to and interactions with the northbound admin API of ONOS-E2T
type E2AdminSession struct {
	E2TEndpoint string
}

// NewSession creates a new admin southbound session of ONOS-KPIMON
func NewSession(e2tEndpoint string) *E2AdminSession {
	log.Info("Creating RicAPIAdminSession")
	return &E2AdminSession{
		E2TEndpoint: e2tEndpoint,
	}
}

// GetListE2NodeIDs returns the list of E2NodeIDs which are connected to ONOS-RIC
func (s *E2AdminSession) GetListE2NodeIDs() ([]string, error) {
	var nodeIDs []string

	adminClient, err := s.connectionHandler()
	if err != nil {
		return []string{}, err
	}

	e2NodeIDStream, err := adminClient.ListE2NodeConnections(context.Background(), &adminapi.ListE2NodeConnectionsRequest{})
	if err != nil {
		log.Errorf("Failed to call ListE2NodeConnections")
		return []string{}, err
	}

	for {
		e2NodeIDStream, err := e2NodeIDStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Errorf("Failed to get e2NodeID")
			return []string{}, err
		} else if e2NodeIDStream != nil {
			nodeIDs = append(nodeIDs, e2NodeIDStream.Id)
		}
	}

	return nodeIDs, nil
}

// GetRANFunctions get list of RAN functions for a given node
func (s *E2AdminSession) GetRANFunctions(nodeID string) ([]*adminapi.RANFunction, error) {
	var ranFunctions []*adminapi.RANFunction

	adminClient, err := s.connectionHandler()
	if err != nil {
		return nil, err
	}
	connections, err := adminClient.ListE2NodeConnections(context.Background(), &adminapi.ListE2NodeConnectionsRequest{})

	if err != nil {
		return nil, err
	}

	for {
		connection, err := connections.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if connection != nil {
			if connection.Id == nodeID {
				ranFunctions = connection.RanFunctions
			}
		}
	}
	return ranFunctions, nil
}

func (s *E2AdminSession) connectionHandler() (adminapi.E2TAdminServiceClient, error) {
	log.Infof("Connecting to ONOS-E2T ... %s", s.E2TEndpoint)

	opts := []grpc.DialOption{
		grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(100 * time.Microsecond)),
	}

	conn, err := southbound.Connect(context.Background(), s.E2TEndpoint, "", "", opts...)
	if err != nil {
		log.Errorf("Failed to connect: %s", err)
		return nil, err
	}

	log.Infof("Connected to %s", s.E2TEndpoint)

	adminClient := adminapi.NewE2TAdminServiceClient(conn)
	return adminClient, nil

}
