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
	res, err := ph.data.GetProducts()
	if err != nil {
		ph.log.WithError(err).Error("Unable to get products")
		// Unexpected error, do not show details to the user
		genericErrorResponse(rw, http.StatusInternalServerError, "Products are not available.")
		return
	}

	// serialize list of products to JSON
	err = utils.ToJSON(res, rw)
	if err != nil {
		ph.log.WithError(err).Error("Unable to serialize to JSON")
		genericErrorResponse(rw, http.StatusInternalServerError, "Products are not available.")
		return
	}
}

func (ph *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p := &data.Product{}

	err := utils.FromJSON(p, r.Body)
	if err != nil {
		ph.log.WithError(err).Error("Unable to deserialize from JSON")
		genericErrorResponse(rw, http.StatusBadRequest, "Product has invalid structure.")
		return
	}

	ph.data.AddProduct(p)
}
