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

// InMemoryProductData implements data.ProductData interface
type InMemoryProductData struct{}

func (d *InMemoryProductData) GetProducts() data.Products {
	return data.Products{&data.Product{ID: 4}}
}
func (d *InMemoryProductData) AddProduct(p *data.Product) {}

func TestGetProductsAsValidJSON(t *testing.T) {
	ds := &InMemoryProductData{}

	ih := NewProduct(logrus.WithField("context", "tests"), ds)

	r := httptest.NewRequest(http.MethodGet, "/products", nil)
	response := httptest.NewRecorder()
	ih.ServeHTTP(response, r)

	assert.Equal(t, response.Code, http.StatusOK, "status should be 200")

	p := make(data.Products, 0, 1)
	err := utils.FromJSON(&p, response.Body)
	assert.NoError(t, err)

	assert.Len(t, p, 1)
	assert.Equal(t, p[0].ID, 4)
}
