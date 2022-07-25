package data

import (
	"context"

	"github.com/sirupsen/logrus"
	protos "github.com/zenichi/green-api/pricing-service/pkg/protos/rates"
)

type ProductData interface {
	GetProducts(ctx context.Context, currency string) (Products, error)
	GetProductById(ctx context.Context, id int, currency string) (*Product, error)
	AddProduct(p *Product) error
	UpdateProduct(p *Product) error
	DeleteProduct(id int) error
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

// GetProducts fetches and returns all products from data store
func (db *ProductDB) GetProducts(ctx context.Context, currency string) (Products, error) {
	log := db.log.WithField("layer", "data")
	log.Info("get products")

	if currency == "" || currency == BaseProductCurrency {
		return internalDB, nil
	}

	rate, err := db.getRate(ctx, currency)
	if err != nil {
		db.log.WithError(err).Info("can not get rates")
		return nil, err
	}

	pl := Products{}
	for _, p := range internalDB {
		pc := *p
		pc.Price = pc.Price * rate
		pl = append(pl, &pc)
	}

	return pl, nil
}

func (db *ProductDB) GetProductById(ctx context.Context, id int, currency string) (*Product, error) {
	i := findIndexByID(id)
	if i < 0 {
		return nil, ErrProductNotFound
	}

	if currency == "" || currency == BaseProductCurrency {
		return internalDB[i], nil
	}

	rate, err := db.getRate(ctx, currency)
	if err != nil {
		db.log.WithError(err).Info("can not get rates")
		return nil, err
	}

	p := *internalDB[i]
	p.Price = p.Price * rate

	return &p, nil
}

// AddProduct adds a product to the datastore
func (db *ProductDB) AddProduct(p *Product) error {
	p.ID = getNextId()
	internalDB = append(internalDB, p)
	return nil
}

// UpdateProduct replace a product in the datastore with the given item
func (db *ProductDB) UpdateProduct(p *Product) error {
	i := findIndexByID(p.ID)
	if i < 0 {
		return ErrProductNotFound
	}

	internalDB[i] = p

	return nil
}

// DeleteProduct removes product from data store with the given ID
func (db *ProductDB) DeleteProduct(id int) error {
	i := findIndexByID(id)
	if i < 0 {
		return ErrProductNotFound
	}

	internalDB = append(internalDB[:i], internalDB[i+1:]...)

	return nil
}

func getNextId() int {
	last := internalDB[len(internalDB)-1]
	return last.ID + 1
}

func findIndexByID(id int) int {
	for i, p := range internalDB {
		if p.ID == id {
			return i
		}
	}

	return -1
}

func (db *ProductDB) getRate(ctx context.Context, toCurrency string) (float64, error) {
	d := &protos.RateRequest{
		FromCurrency: BaseProductCurrency,
		ToCurrency:   toCurrency,
	}
	rr, err := db.rates.GetRate(ctx, d)
	if err != nil {
		return 0, err
	}

	return rr.Rate, nil
}

var internalDB = Products{
	&Product{
		ID:          1,
		Name:        "Garlic",
		Description: "Spicy Chinese garlic",
		Price:       1.50,
		ExternalID:  "G-45646",
	},
	&Product{
		ID:          2,
		Name:        "Onion",
		Description: "Small yellow onion",
		Price:       1.1,
		ExternalID:  "O-45234",
	},
}
