package currency

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func (cp CurrencyProvider) GetRate(base string) (map[string]float32, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	spaceClient := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 2, // Timeout after 2 seconds
	}
	url := cp.url
	if base != "" {
		url = url + "?base=" + base
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
		return nil, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
		return nil, getErr
	}

	reqFormat := APIRequest{}
	jsonErr := json.Unmarshal(body, &reqFormat)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return reqFormat.Rates, nil
}
