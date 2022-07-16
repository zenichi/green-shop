package data

import (
	"context"

	"github.com/sirupsen/logrus"
	protos "github.com/zenichi/green-api/pricing-service/pkg/protos/rates"
)

type ProductData interface {
	GetProducts() (Products, error)
	AddProduct(p *Product)
}

// ProductDB defines methods available on product data
type ProductDB struct {
	log   *logrus.Entry
	rates protos.RateServiceClient
}

// NewProductDB creates the DB access handler
func NewProductDB(log *logrus.Entry, rates protos.RateServiceClient) *ProductDB {
	return &ProductDB{log, rates}
}

func (db *ProductDB) GetProducts() (Products, error) {
	log := db.log.WithField("layer", "data")
	log.Info("get products")

	ctx := context.Background()
	d := &protos.RateRequest{
		FromCurrency: "USD",
		ToCurrency:   "EUR",
	}
	rr, err := db.rates.GetRate(ctx, d)
	if err != nil {
		log.WithError(err).Info("can not get rates")
		return nil, err
	}

	pl := Products{}
	for _, p := range internalDB {
		pc := *p
		pc.Price = pc.Price * rr.Rate
		pl = append(pl, &pc)
	}

	return pl, nil
}

func (db *ProductDB) AddProduct(p *Product) {
	p.ID = getNextId()
	internalDB = append(internalDB, p)
}

func getNextId() int {
	last := internalDB[len(internalDB)-1]
	return last.ID + 1
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
