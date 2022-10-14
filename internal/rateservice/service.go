package rateservice

import (
	"context"
	"fmt"
	"strings"

	protos "github.com/mannanmcc/proto/rates/rate"
	errors "github.com/mannanmcc/rateService/internal/errors"
)

type Service struct {
	protos.UnimplementedRateServiceServer
}

var currentRate = map[string]string{
	"gbp-gbp": "1",
	"gbp-usd": "1.1",
	"usd-gbp": "0.9",
	"usd-usd": "1",
}

const invalidRequest = errors.Error("invalid request")
const currencyNotSupported = errors.Error("un supported currency provided")

func New() *Service {
	return &Service{}
}

func (s *Service) GetRate(ctx context.Context, req *protos.RateRequest) (*protos.RateResponse, error) {
	if req.GetBaseCurrency() == "" || req.GetTargetCurrency() == "" {
		return nil, invalidRequest
	}

	var rate string
	var ok bool

	combinationMatch := fmt.Sprintf("%s-%s", strings.ToLower(req.GetBaseCurrency()), strings.ToLower(req.GetTargetCurrency()))
	if rate, ok = currentRate[combinationMatch]; !ok {
		return nil, currencyNotSupported
	}

	return &protos.RateResponse{Rate: rate}, nil
}
