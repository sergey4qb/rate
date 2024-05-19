package app

import (
	"github.com/sergey4qb/rate/internal/core/ports"
	"github.com/sergey4qb/rate/internal/core/service/exchange_rates_service"
	"github.com/sergey4qb/rate/internal/core/service/notification_service"
	"github.com/sergey4qb/rate/internal/core/service/rates_subscribers_service"
	"os"
)

type services struct {
	exchangeRateService     ports.ExchangeRatesService
	ratesSubscriberService  ports.RatesSubscribersService
	rateNotificationService ports.RateNotificationService
}

func createServices(infrastructure *infrastructure) *services {
	exchangeRateService := exchange_rates_service.NewService(infrastructure.exchangeRateClient)
	ratesSubscriberService := rates_subscribers_service.NewService(infrastructure.ratesSubscriberStorage)
	config := notification_service.Config{
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
		From:     os.Getenv("SMTP_FROM"),
		Subject:  os.Getenv("SMTP_SUBJECT"),
	}
	notificationService := notification_service.NewService(
		&config,
		ratesSubscriberService,
		exchangeRateService,
	)
	return &services{
		exchangeRateService:     exchangeRateService,
		ratesSubscriberService:  ratesSubscriberService,
		rateNotificationService: notificationService,
	}
}
