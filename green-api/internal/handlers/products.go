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
	v    *data.Validator
}

// NewProduct creates the new Product handler with the given logger and db access
func NewProduct(log *logrus.Entry, data data.ProductData, v *data.Validator) *Product {
	return &Product{log, data, v}
}

func (ph *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
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

	rw.WriteHeader(http.StatusOK)
}

func (ph *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p := &data.Product{}

	err := utils.FromJSON(p, r.Body)
	if err != nil {
		ph.log.WithError(err).Error("Unable to deserialize from JSON")
		genericErrorResponse(rw, http.StatusBadRequest, "Product has invalid structure.")
		return
	}

	errors := ph.v.Validate(p)
	if len(errors) > 0 {
		validationErrorsResponse(rw, errors)
		return
	}

	ph.data.AddProduct(p)
	rw.WriteHeader(http.StatusOK)
}
