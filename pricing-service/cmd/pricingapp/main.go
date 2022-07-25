package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/zenichi/green-api/pricing-service/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "github.com/zenichi/green-api/pricing-service/pkg/protos/rates"
)

var serverAddress = flag.String("pricing-service-addr", "localhost:9085", "the address for the server")

func main() {
	flag.Parse()

	log := logrus.WithField("app", "pricing-service")
	log.WithField("addr", *serverAddress).Info("initializing service")

	// initializing new grpc server with insecure to allow http connections (on localhost)
	gs := grpc.NewServer()

	// create an instance of rates server
	rs := server.NewRates(log)

	// register rate service server
	protos.RegisterRateServiceServer(gs, rs)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// create a tcp socket for inbound server connections
	l, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.WithError(err).Error("Unable to create listener")
		os.Exit(1)
	}

	go func() {
		err = gs.Serve(l)
		if err != nil {
			log.WithError(err).Error("Listener shutted down")
			os.Exit(1)
		}
	}()

	idleConnsClosed := make(chan os.Signal, 1)
	signal.Notify(idleConnsClosed, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	// wait until signal received
	sig := <-idleConnsClosed
	log.WithField("signal", sig).Info("gracefully shutting down server")

	gs.GracefulStop()
	log.Info("server shutted down")
}
