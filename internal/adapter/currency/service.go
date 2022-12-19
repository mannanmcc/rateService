package currency

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mannanmcc/rateService/internal/logger"
	"go.uber.org/zap"
)

type CurrencyProvider struct {
	url string
}

type APIRequest struct {
	Base    string             `json:"base"`
	Success bool               `json:"success"`
	Date    string             `json:"date"`
	Rates   map[string]float32 `json:"rates"`
}

func New(url string) CurrencyProvider {
	return CurrencyProvider{
		url: url,
	}
}

func (cp CurrencyProvider) GetRate(ctx context.Context, base string) (map[string]float32, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	spaceClient := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 2, // Timeout after 2 seconds
	}
	url := cp.url
	if base != "" {
		url = url + "?base=" + strings.ToLower(base)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error(ctx, "failed to initialize http request", zap.Error(err))
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		logger.Print(ctx, "failed to get rate from remote rate service", zap.Error(getErr))
		return nil, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		logger.Print(ctx, "could not read  response payload", zap.Error(getErr))
		return nil, getErr
	}

	reqFormat := APIRequest{}
	jsonErr := json.Unmarshal(body, &reqFormat)
	if jsonErr != nil {
		logger.Print(ctx, "could not unmarshall response payload", zap.Error(jsonErr))
	}

	return reqFormat.Rates, nil
}
