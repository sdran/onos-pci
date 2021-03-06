// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"

	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-pci/pkg/manager"
)

var log = logging.GetLogger("main")

func main() {
	caPath := flag.String("caPath", "", "path to CA certificate")
	keyPath := flag.String("keyPath", "", "path to client private key")
	certPath := flag.String("certPath", "", "path to client certificate")
	e2tEndpoint := flag.String("e2tEndpoint", "onos-e2t:5150", "E2T service endpoint")
	e2subEndpoint := flag.String("e2subEndpoint", "onos-e2sub:5150", "E2Sub service endpoint")
	ricActionID := flag.Int("ricActionID", 10, "RIC Action ID in E2 message")
	grpcPort := flag.Int("grpcPort", 5150, "grpc Port number")

	ready := make(chan bool)

	flag.Parse()

	_, err := certs.HandleCertPaths(*caPath, *keyPath, *certPath, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Starting onos-pci")

	cfg := manager.Config{
		CAPath:        *caPath,
		KeyPath:       *keyPath,
		CertPath:      *certPath,
		E2tEndpoint:   *e2tEndpoint,
		E2SubEndpoint: *e2subEndpoint,
		GRPCPort:      *grpcPort,
		RicActionID:   int32(*ricActionID),
	}

	mgr := manager.NewManager(cfg)
	mgr.Run()
	<-ready
}
