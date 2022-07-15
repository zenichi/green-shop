package data

// Product defines structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	ExternalID  string  `json:"externalId"`
}

// Products represents a collection of Product items
type Products []*Product
