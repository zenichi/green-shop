package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/zenichi/green-shop/green-api/internal/data"
	"github.com/zenichi/green-shop/green-api/internal/utils"
)

var (
	SimpleGetProductsRequest = httptest.NewRequest("GET", "/products", nil)
)

func runRequest(t *testing.T, srv http.Handler, r *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	srv.ServeHTTP(response, r)

	return response
}

// InMemoryProductData implements data.ProductData interface
type InMemoryProductData struct{}

func (d *InMemoryProductData) GetProducts() data.Products {
	return data.Products{&data.Product{ID: 4}}
}
func (d *InMemoryProductData) AddProduct(p *data.Product) {}

func TestGetProductsAsValidJSON(t *testing.T) {
	// create dummy store
	ds := &InMemoryProductData{}

	// create handler
	ph := NewProduct(logrus.WithField("context", "tests"), ds)

	// run request
	response := runRequest(t, ph, SimpleGetProductsRequest)

	// asserts response
	assert.Equal(t, response.Code, http.StatusOK, "status should be 200")

	var p data.Products
	err := utils.FromJSON(&p, response.Body)
	assert.NoError(t, err)

	assert.Len(t, p, 1)
	assert.Equal(t, p[0].ID, 4)
}
