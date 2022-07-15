package data

import "github.com/sirupsen/logrus"

// ProductDB defines methods available on product data
type ProductDB struct {
	log *logrus.Entry
}

// NewProductDB creates the DB access handler
func NewProductDB(log *logrus.Entry) *ProductDB {
	return &ProductDB{log}
}

func (db *ProductDB) GetProducts() Products {
	log := db.log.WithField("layer", "data")
	log.Info("get products")

	return internalDB
}

var internalDB = Products{
	&Product{
		ID:          1,
		Name:        "Garlic",
		Description: "Spicy Chinese garlic",
		Price:       1.50,
		Currency:    "USD",
		ExternalID:  "G-45646",
	},
	&Product{
		ID:          2,
		Name:        "Onion",
		Description: "Small yellow onion",
		Price:       1.1,
		Currency:    "USD",
		ExternalID:  "O-45234",
	},
}
