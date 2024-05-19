package ports

import (
	"context"
	"github.com/sergey4qb/rate/internal/infrastructure/external/http/exchange_rates_client"
)

type ExchangeRatesClient interface {
	GetUAHUSDExchangeRate(ctx context.Context) (*exchange_rates_client.Response, error)
}

type ExchangeRatesService interface {
	GetUAHUSDExchangeRate(ctx context.Context) (float64, error)
}
