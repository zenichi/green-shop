package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// GetProducts handles GET requests to list all products
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

// GetProduct handles GET requests to list a signle product
func (ph *Product) GetSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductIdParam(r)

	// fetch the product from the datastore
	p, err := ph.data.GetProductById(id)
	if err != nil {
		if err == data.ErrProductNotFound {
			genericErrorResponse(rw, http.StatusNotFound, "Product not found in the database")
		} else {
			ph.log.WithError(err).Error("Unable to get product")
			genericErrorResponse(rw, http.StatusInternalServerError, "Product can not be retrieved.")
		}
		return
	}

	// serialize product to JSON
	err = utils.ToJSON(p, rw)
	if err != nil {
		ph.log.WithError(err).Error("Unable to serialize to JSON")
		genericErrorResponse(rw, http.StatusInternalServerError, "Product is not available.")
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// AddProduct handles POST requests to add new products
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

// UpdateProduct handles PUT requests to update products
func (ph *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
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

	err = ph.data.UpdateProduct(p)
	if err != nil {
		if err == data.ErrProductNotFound {
			genericErrorResponse(rw, http.StatusNotFound, "Product not found in the database")
		} else {
			ph.log.WithError(err).Error("Unable to update product")
			genericErrorResponse(rw, http.StatusInternalServerError, "Product can not be updated.")
		}
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// DeleteProducs handles HttpDelete requests to remove product by given ID param
func (ph *Product) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductIdParam(r)

	err := ph.data.DeleteProduct(id)
	if err != nil {
		if err == data.ErrProductNotFound {
			genericErrorResponse(rw, http.StatusNotFound, "Product not found in the database")
		} else {
			ph.log.WithError(err).Error("Unable to delete product")
			genericErrorResponse(rw, http.StatusInternalServerError, "Product can not be updated.")
		}
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// getProductIdParam reads product ID URI param and parses it to int
func getProductIdParam(r *http.Request) int {
	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// should never happen as router ensures param is valid
		log.Fatalf("invalid id: %v", id)
	}
	return id
}
