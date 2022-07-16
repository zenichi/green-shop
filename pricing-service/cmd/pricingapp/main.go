package main

import (
	"flag"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var serverAddress = flag.String("pricing-service-addr", "localhost:9085", "the address for the server")

func main() {
	flag.Parse()

	log := logrus.WithField("app", "pricing-service")
	log.WithField("addr", *serverAddress).Info("initializing service")

	// initializing new grpc server with insecure to allow http connections (on localhost)
	gs := grpc.NewServer()

	// create a tcp socket for inbound server connections
	l, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.WithError(err).Error("Unable to create listener")
		os.Exit(1)
	}

	err = gs.Serve(l)
	if err != nil {
		log.WithError(err).Error("Listener shutted down")
	}
}
