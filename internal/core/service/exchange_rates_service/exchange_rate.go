package exchange_rates_service

import (
	"context"
	"github.com/sergey4qb/rate/internal/core/ports"
)

type Service struct {
	client ports.ExchangeRatesClient
}

func NewService(client ports.ExchangeRatesClient) *Service {
	return &Service{client: client}
}

func (s *Service) GetUAHUSDExchangeRate(ctx context.Context) (float64, error) {
	rate, err := s.client.GetUAHUSDExchangeRate(ctx)
	if err != nil {
		return 0, err
	}
	return rate.Rate, nil
}
