package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/zenichi/green-shop/green-api/internal/data"
)

// Product is a handler for products REST API
type Product struct {
	log  *logrus.Entry
	data *data.ProductDB
}

// NewProduct creates the new Product handler with the given logger and db access
func NewProduct(log *logrus.Entry, data *data.ProductDB) *Product {
	return &Product{log, data}
}

func (ph *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ph.getProducts(rw, r)
		return
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
	rw.Header().Set("Allow", "GET")
	rw.Write([]byte("Method is not allowed"))
}

func (ph *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	// fetch the products from the datastore
	res := ph.data.GetProducts()

	// serialize list of products to JSON
	e := json.NewEncoder(rw)
	err := e.Encode(res)
	if err != nil {
		ph.log.WithError(err).Error("Unable to serialize to JSON")
		http.Error(rw, "Unable to serialize products", http.StatusInternalServerError)
	}
}
