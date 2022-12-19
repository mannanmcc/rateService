package rateservice

import (
	"context"
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

const ErrInvalidRequest = errors.Error("invalid request")
const ErrApiCallFailed = errors.Error("failed to retrieve rate fro remote api")
const ErrCurrencyNotSupported = errors.Error("un supported currency provided")
const errFailedToGetCurrency = errors.Error("Failed to get currency")

func New(provider currency.CurrencyProvider) *Service {
	return &Service{
		currencyProvider: provider,
	}
}

func (s *Service) GetRate(ctx context.Context, req Request) (Response, error) {

	response := Response{}
	var err error
	var rates map[string]float32

	if !isCurrencySupported(strings.ToUpper(req.TargetCurrency)) {
		return response, ErrCurrencyNotSupported
	}

	if rates, err = s.currencyProvider.GetRate(req.BaseCurrency); err != nil {
		log.Println("failed to retrieve rate from remote rate provide")
		return response, ErrApiCallFailed
	}

	if rateFromApi, ok := rates[strings.ToUpper(req.TargetCurrency)]; ok {
		return Response{Rate: rateFromApi}, nil
	}

	return response, errFailedToGetCurrency
}

func isCurrencySupported(curr string) bool {
	for _, currency := range supportedCurrencies {
		if currency == curr {
			return true
		}
	}

	return false
}
