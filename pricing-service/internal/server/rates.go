package server

import (
	context "context"

	"github.com/sirupsen/logrus"
	protos "github.com/zenichi/green-api/pricing-service/pkg/protos/rates"
)

// Rates is a grpc server, it implements methods defined by RateServiceServer
type Rates struct {
	protos.UnimplementedRateServiceServer
	log *logrus.Entry
}

func NewRates(log *logrus.Entry) *Rates {
	return &Rates{log: log}
}

// GetRate implements rates.RateServiceServer
func (c *Rates) GetRate(ctx context.Context, r *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.WithFields(logrus.Fields{"From": r.FromCurrency, "To": r.ToCurrency}).Info("handle request for GetRate")

	// todo: get rates dynamically from external server
	var rate float64
	switch r.ToCurrency {
	case "EUR":
		rate = 1.1
	case "GBP":
		rate = 1.3
	default:
		rate = 0.8
	}

	return &protos.RateResponse{Rate: rate}, nil
}
