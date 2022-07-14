package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zenichi/green-shop/green-api/internal/handlers"
)

var serverAddress = flag.String("green-api-addr", "localhost:9081", "the address for the server to listen on, in the form 'host:port'")

func main() {
	flag.Parse()

	log := logrus.WithField("app", "green-api")
	log.WithField("addr", *serverAddress).Info("initializing server")

	// create the handlers
	ih := handlers.NewInfo(log)

	// create the new ServeMux and register handler
	mux := http.NewServeMux()
	mux.Handle("/info", ih)

	// create the HttpServer
	srv := http.Server{
		Addr:         *serverAddress,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second, // max time for connections Keep-Alive
	}

	// run HttpServer
	err := srv.ListenAndServe()
	if err != nil {
		log.WithError(err).Error("server stopped")
	}
}
