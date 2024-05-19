package exchange_rates_client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeRateClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewExchangeRatePublicClient(baseURL string) *ExchangeRateClient {
	return &ExchangeRateClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *ExchangeRateClient) GetUAHUSDExchangeRate(ctx context.Context) (*Response, error) {
	resp, err := c.HTTPClient.Get(c.BaseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var exchangeRateResponse []Response
	if err := json.NewDecoder(resp.Body).Decode(&exchangeRateResponse); err != nil {
		return nil, err
	}

	return &exchangeRateResponse[0], nil
}
