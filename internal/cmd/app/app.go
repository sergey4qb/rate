package app

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

const url = "https://bank.gov.ua/NBUStatService/v1/statdirectory/dollar_info?json"

type App struct {
	storage        *storage
	infrastructure *infrastructure
	services       *services
	delivery       *delivery
}

func Create() *App {
	storage := createStorage()
	infrastructure := createInfrastructure(storage)
	services := createServices(infrastructure)
	delivery := createDelivery(services)
	return &App{
		services:       services,
		storage:        storage,
		infrastructure: infrastructure,
		delivery:       delivery,
	}
}

func (a *App) Start(ctx context.Context) error {
	a.delivery.Start()
	scheduledTime := os.Getenv("SCHEDULED_TIME")
	errChan := make(chan error)
	log.Println("Notification will be sent everyday at:", scheduledTime)
	go a.handleSendErrors(errChan)
	err := a.runScheduledSending(
		ctx,
		scheduledTime,
		a.services.rateNotificationService.SendRates,
		errChan,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) runScheduledSending(
	ctx context.Context,
	scheduleTime string,
	send func(ctx context.Context) error,
	errChan chan<- error,
) error {
	t, err := time.Parse("15:04", scheduleTime)
	if err != nil {
		return fmt.Errorf("invalid time format: %w", err)
	}
	cronExpr := fmt.Sprintf("%d %d * * *", t.Minute(), t.Hour())
	c := cron.New()
	_, err = c.AddFunc(cronExpr, func() {
		if e := send(ctx); e != nil {
			errChan <- e
		}
	})
	if err != nil {
		return fmt.Errorf("error scheduling email sending: %w", err)
	}
	c.Start()
	return nil
}

func (a *App) handleSendErrors(errChan <-chan error) {
	for {
		select {
		case e := <-errChan:
			log.Println("error while sending notification:", e)
		}
	}
}
