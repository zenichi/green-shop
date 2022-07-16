package data

// Product defines structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Currency    string  `json:"currency" validate:"required,len=3"`
	ExternalID  string  `json:"externalId" validate:"required"`
}

// Products represents a collection of Product items
type Products []*Product
