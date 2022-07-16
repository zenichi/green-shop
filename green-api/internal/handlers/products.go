package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/zenichi/green-shop/green-api/internal/data"
	"github.com/zenichi/green-shop/green-api/internal/utils"
)

// Product is a handler for products REST API
type Product struct {
	log  *logrus.Entry
	data data.ProductData
}

// NewProduct creates the new Product handler with the given logger and db access
func NewProduct(log *logrus.Entry, data data.ProductData) *Product {
	return &Product{log, data}
}

func (ph *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ph.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		ph.addProduct(rw, r)
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
	err := utils.ToJSON(res, rw)
	if err != nil {
		ph.log.WithError(err).Error("Unable to serialize to JSON")
		http.Error(rw, "Unable to serialize products", http.StatusInternalServerError)
	}
}

func (ph *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p := &data.Product{}

	err := utils.FromJSON(p, r.Body)
	if err != nil {
		ph.log.WithError(err).Error("Unable to deserialize from JSON")
		http.Error(rw, "Unable to deserialize product", http.StatusBadRequest)
	}

	ph.data.AddProduct(p)
}
