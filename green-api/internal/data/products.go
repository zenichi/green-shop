package data

import "errors"

var (
	ErrProductNotFound  = errors.New("Product not found")
	BaseProductCurrency = "USD"
)

// Product defines structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	ExternalID  string  `json:"externalId" validate:"required"`
}

// Products represents a collection of Product items
type Products []*Product
