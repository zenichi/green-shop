package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsError(t *testing.T) {
	p := &Product{
		ID:          1,
		Name:        "",
		Description: "Spicy Chinese garlic",
		Price:       1.50,
		ExternalID:  "G-45646",
	}

	v := NewValidator()
	errors := v.Validate(p)

	assert.Len(t, errors, 1)
	assert.Contains(t, errors[0], "Product.Name")
}

func TestProductPriceLowerThanZeroReturnsError(t *testing.T) {
	p := &Product{
		ID:          1,
		Name:        "Garlic",
		Description: "Spicy Chinese garlic",
		Price:       -1.50,
		ExternalID:  "G-45646",
	}

	v := NewValidator()
	errors := v.Validate(p)

	assert.Len(t, errors, 1)
	assert.Contains(t, errors[0], "Product.Price")
}
