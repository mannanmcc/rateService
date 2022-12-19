package rateservice

type Request struct {
	BaseCurrency   string
	TargetCurrency string
}

type Response struct {
	Rate float32
}
