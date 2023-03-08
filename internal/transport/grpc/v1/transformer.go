package v1

import (
	"fmt"

	protos "github.com/mannanmcc/proto/rates/rate"

	"github.com/mannanmcc/rateService/internal/rateservice"
)

func transformRequest(req *protos.RateRequest) rateservice.Request {
	return rateservice.Request{
		BaseCurrency:   req.GetBaseCurrency(),
		TargetCurrency: req.GetTargetCurrency(),
	}
}

func transformResponse(response rateservice.Response) *protos.RateResponse {
	return &protos.RateResponse{
		Rate: fmt.Sprintf("%.2f", response.Rate),
	}
}
