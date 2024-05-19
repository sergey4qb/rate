package app

import (
	"github.com/sergey4qb/rate/internal/core/ports"
	"github.com/sergey4qb/rate/internal/infrastructure/external/http/exchange_rates_client"
	"github.com/sergey4qb/rate/internal/infrastructure/storage/rates_subscribers"
)

type infrastructure struct {
	exchangeRateClient     ports.ExchangeRatesClient
	ratesSubscriberStorage ports.RatesSubscribersStorage
}

func createInfrastructure(storage *storage) *infrastructure {
	return &infrastructure{
		exchangeRateClient:     exchange_rates_client.NewExchangeRatePublicClient(url),
		ratesSubscriberStorage: rates_subscribers.NewStorage(storage.db).Db,
	}
}
