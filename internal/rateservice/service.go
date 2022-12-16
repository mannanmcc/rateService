package rateservice

import (
	"context"
	"fmt"
	"log"
	"strings"

	protos "github.com/mannanmcc/proto/rates/rate"
	currency "github.com/mannanmcc/rateService/internal/adapter/currency"
	errors "github.com/mannanmcc/rateService/internal/errors"
)

type Service struct {
	protos.UnimplementedRateServiceServer
	currencyProvider currency.CurrencyProvider
}

var supportedCurrencies = []string{"USD", "EUR", "GBP", "BDT"}

const invalidRequest = errors.Error("invalid request")
const apiCallFailedErr = errors.Error("failed to retrieve rate fro remote api")
const currencyNotSupported = errors.Error("un supported currency provided")
const errFailedToGetCurrency = errors.Error("Failed to get currency")

func New(provider currency.CurrencyProvider) *Service {
	return &Service{
		currencyProvider: provider,
	}
}

func (s *Service) GetRate(ctx context.Context, req *protos.RateRequest) (*protos.RateResponse, error) {
	if req.GetBaseCurrency() == "" || req.GetTargetCurrency() == "" {
		return nil, invalidRequest
	}

	var err error
	var rates map[string]float32

	currencyForRate := strings.ToUpper(req.GetTargetCurrency())

	if !isCurrenctSupported(currencyForRate) {
		return nil, currencyNotSupported
	}

	if rates, err = s.currencyProvider.GetRate(req.GetBaseCurrency()); err != nil {
		log.Fatal(err)
		return nil, apiCallFailedErr
	}

	if rateFromApi, ok := rates[strings.ToUpper(req.GetTargetCurrency())]; ok {
		return &protos.RateResponse{Rate: fmt.Sprintf("%.2f", rateFromApi)}, nil
	}

	return nil, errFailedToGetCurrency
}

func isCurrenctSupported(curr string) bool {
	for _, currency := range supportedCurrencies {
		if currency == curr {
			return true
		}
	}

	return false
}
