package rateservice

import (
	"context"
	"fmt"
	"strings"

	protos "github.com/mannanmcc/proto/rates/rate"
	currency "github.com/mannanmcc/rateService/internal/adapter/currency"
	errors "github.com/mannanmcc/rateService/internal/errors"
	"github.com/mannanmcc/rateService/internal/logger"
	"github.com/mannanmcc/rateService/store"
	"go.uber.org/zap"
)

type Service struct {
	store            *store.Store
	currencyProvider currency.Provider
	protos.UnimplementedRateServiceServer
}

var supportedCurrencies = []string{"USD", "EUR", "GBP", "BDT"}

const ErrInvalidRequest = errors.Error("invalid request")
const ErrAPICallFailed = errors.Error("failed to retrieve rate fro remote api")
const ErrCurrencyNotSupported = errors.Error("un supported currency provided")
const errFailedToGetCurrency = errors.Error("Failed to get currency")

func New(provider currency.Provider, st *store.Store) *Service {
	return &Service{
		currencyProvider: provider,
		store:            st,
	}
}

func (s *Service) GetRate(ctx context.Context, req Request) (Response, error) {
	var err error
	var rates map[string]float32

	response := Response{}

	if err = req.validate(); err != nil {
		return response, ErrInvalidRequest.Wrap(err)
	}

	if !isCurrencySupported(strings.ToUpper(req.TargetCurrency)) {
		return response, ErrCurrencyNotSupported
	}

	rateKey := fmt.Sprintf("%s_%s", strings.ToUpper(req.BaseCurrency), strings.ToUpper(req.TargetCurrency))
	//check if rate exists in the db
	logger.Print(ctx, "checking rate in db", zap.String("rate_key::", rateKey))
	rate, err := s.store.GetRate(ctx, rateKey)
	if err == nil {
		logger.Print(ctx, "returning rate from db", zap.Float32("rate", rate.RateValue))
		return Response{Rate: rate.RateValue}, nil
	} else {
		logger.Print(ctx, "rate not found ::::::", zap.Error(err))
	}

	if rates, err = s.currencyProvider.GetRate(ctx, req.BaseCurrency); err != nil {
		return response, ErrAPICallFailed.Wrap(err)
	}

	if rateFromAPI, ok := rates[strings.ToUpper(req.TargetCurrency)]; ok {
		// add rate to the store
		logger.Print(ctx, "add rate to db", zap.String("rate_key::", rateKey))
		_, err := s.store.AddRate(ctx, rateFromAPI, rateKey)
		if err != nil {
			return response, err
		}

		return Response{Rate: rateFromAPI}, nil
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
