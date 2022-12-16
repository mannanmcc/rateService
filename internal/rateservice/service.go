package rateservice

import (
	"context"
	"fmt"
	"log"

	protos "github.com/mannanmcc/proto/rates/rate"
	currency "github.com/mannanmcc/rateService/internal/adapter/currency"
	errors "github.com/mannanmcc/rateService/internal/errors"
)

type Service struct {
	//this should not be here
	protos.UnimplementedRateServiceServer
	currencyProvider currency.CurrencyProvider
}

var currentRate = map[string]string{
	"gbp-gbp": "1",
	"gbp-usd": "1.1",
	"usd-gbp": "0.9",
	"usd-usd": "1",
}

const invalidRequest = errors.Error("invalid request")
const apiCallFailedErr = errors.Error("failed to retrieve rate fro remote api")
const currencyNotSupported = errors.Error("un supported currency provided")

func New(provider currency.CurrencyProvider) *Service {
	return &Service{
		currencyProvider: provider,
	}
}

func (s *Service) GetRate(ctx context.Context, req *protos.RateRequest) (*protos.RateResponse, error) {
	if req.GetBaseCurrency() == "" || req.GetTargetCurrency() == "" {
		return nil, invalidRequest
	}

	var ok bool
	var err error
	var rates map[string]float32
	var rateFromApi float32

	if rates, err = s.currencyProvider.GetRate(req.GetBaseCurrency()); err != nil {
		log.Fatal(err)
		return nil, apiCallFailedErr
	}

	if rateFromApi, ok = rates[req.GetTargetCurrency()]; !ok {
		return nil, currencyNotSupported
	}
	//combinationMatch := fmt.Sprintf("%s-%s", strings.ToLower(req.GetBaseCurrency()), strings.ToLower(req.GetTargetCurrency()))
	// if rate, ok = currentRate[combinationMatch]; !ok {
	// 	return nil, currencyNotSupported
	// }

	return &protos.RateResponse{Rate: fmt.Sprintf("%.2f", rateFromApi)}, nil
}
