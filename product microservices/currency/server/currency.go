package server

import (
	"context"

	"github.com/Kr-Harshit/golang-example/product-microservices/currency/data"
	protos "github.com/Kr-Harshit/golang-example/product-microservices/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log   hclog.Logger
	rates *data.ExchangeRates
}

func NewCurrency(l hclog.Logger, r *data.ExchangeRates) *Currency {
	return &Currency{log: l, rates: r}
}

func (c *Currency) GetRate(ctx context.Context, req *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", req.GetBase(), "destination", req.GetDestination())

	rate, err := c.rates.GetRate(req.GetBase().String(), req.GetDestination().String())

	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}
