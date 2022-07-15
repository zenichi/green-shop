package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
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
	m := handlers.NewApiMiddleware(log)
	ih := handlers.NewInfo(log)

	// create the new ServeMux and register handler
	mux := http.NewServeMux()
	mux.Handle("/info", m.WithLogging((ih)))

	// create the HttpServer
	srv := http.Server{
		Addr:         *serverAddress,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second, // max time for connections Keep-Alive
	}

	// run HttpServer async
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.WithError(err).Error("server stopped")
			os.Exit(1) // non zero means error
		}
	}()

	// track os signals to gracefuly shutdown the server
	idleConnsClosed := make(chan os.Signal)
	signal.Notify(idleConnsClosed, os.Interrupt)
	signal.Notify(idleConnsClosed, os.Kill)

	// wait until os.signal received
	sig := <-idleConnsClosed
	log.WithField("signal", sig).Info("gracefully shutting down server")

	// wait until current operations complete and shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	srv.Shutdown(ctx)
}
