package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/go-hclog"
)

// ExchangeRates structure holds the different currency
// rate ratio to a given base currency
type ExchangeRates struct {
	log   hclog.Logger
	base  string
	rates map[string]float64
}

// NewRates returns a new ExhangeRates structure
// base is the three digit code of base currency
func NewRates(l hclog.Logger, base string) (*ExchangeRates, error) {
	er := &ExchangeRates{log: l}

	err := er.getRates(base)
	return er, err
}

// GetRate returns ratio of base currency rate to destination currency rate
func (e *ExchangeRates) GetRate(base, dest string) (float64, error) {
	br, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("rate no found for currency %s", base)
	}

	dr, ok := e.rates[dest]
	if !ok {
		return 0, fmt.Errorf("rate no found for currency %s", dest)
	}

	return dr / br, nil
}

// getRates fetchs all currency rates from freecurrencyapi.net
func (e *ExchangeRates) getRates(base string) error {
	url := fmt.Sprintf("https://freecurrencyapi.net/api/v2/latest?apikey=%s&base_currency=%s",
		os.Getenv("CURRENCY_API"), base)

	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("expected error code 200 get %d", res.StatusCode)
	}
	defer res.Body.Close()

	rates := &Rates{}
	if err := json.NewDecoder(res.Body).Decode(rates); err != nil {
		return err
	}

	e.rates = rates.Data
	e.rates[base] = 1
	e.base = rates.Query.Base

	return nil
}

type Rates struct {
	Query RatesQuery         `json:"query"`
	Data  map[string]float64 `json:"data"`
}

type RatesQuery struct {
	Base      string `json:"base_currency"`
	Timestamp int64  `json:"timestamp"`
}
