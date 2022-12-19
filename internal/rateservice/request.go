package rateservice

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Request struct {
	BaseCurrency   string
	TargetCurrency string
}

func (req Request) validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.BaseCurrency, validation.Required),
		validation.Field(&req.BaseCurrency, validation.Required),
	)
}

type Response struct {
	Rate float32
}
