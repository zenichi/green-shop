package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zenichi/green-shop/green-api/internal/data"
	"github.com/zenichi/green-shop/green-api/internal/handlers"
	protos "github.com/zenichi/green-shop/pricing-service/pkg/protos/rates"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddress = flag.String("green-api-addr", "localhost:9081", "the address for the server to listen on, in the form 'host:port'")
var clientAddress = flag.String("green-pricing-client-addr", "localhost:9085", "the address for the grpc server to connect")

func main() {
	flag.Parse()

	log := logrus.WithField("app", "green-api")
	log.WithField("addr", *serverAddress).Info("initializing server")

	// create rates client for grpc server
	conn, err := grpc.Dial(*clientAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	cl := protos.NewRateServiceClient(conn)

	// create the handlers
	m := handlers.NewApiMiddleware(log)
	ih := handlers.NewInfo(log)
	pd := data.NewProductDB(log, cl)
	v := data.NewValidator()
	ph := handlers.NewProduct(log, pd, v)

	// create the new router and register handlers
	mux := mux.NewRouter()
	mux.Use(m.WithLogging)

	getR := mux.Methods(http.MethodGet).Subrouter()
	getR.Handle("/info", ih)
	getR.HandleFunc("/products/{currency:[A-Z]{3}}", ph.GetProducts)
	getR.HandleFunc("/products", ph.GetProducts)
	getR.HandleFunc("/products/{id:[0-9]+}/{currency:[A-Z]{3}}", ph.GetSingle)
	getR.HandleFunc("/products/{id:[0-9]+}", ph.GetSingle)

	mux.Handle("/products", ph.ValidateProduct(http.HandlerFunc(ph.AddProduct))).Methods(http.MethodPost)
	mux.Handle("/products", ph.ValidateProduct(http.HandlerFunc(ph.UpdateProduct))).Methods(http.MethodPut)
	mux.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct).Methods(http.MethodDelete)

	// CORS
	headersOk := gohandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Access-Control-Allow-Origin"})
	originsOk := gohandlers.AllowedOrigins([]string{"*"})
	methodsOk := gohandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	cors := gohandlers.CORS(headersOk, originsOk, methodsOk)

	// create the HttpServer
	srv := http.Server{
		Addr:         *serverAddress,
		Handler:      cors(mux),
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
	idleConnsClosed := make(chan os.Signal, 1)
	signal.Notify(idleConnsClosed, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// wait until os.signal received
	sig := <-idleConnsClosed
	log.WithField("signal", sig).Info("gracefully shutting down server")

	// wait until current operations complete and shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.WithError(err).Error("shutdown error")
	} else {
		log.Info("gracefully stopped")
	}
}
