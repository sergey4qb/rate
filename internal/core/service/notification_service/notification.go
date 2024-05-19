package notification_service

import (
	"context"
	"github.com/sergey4qb/rate/internal/core/domain/subscriber"
	"github.com/sergey4qb/rate/internal/core/ports"
	"net/smtp"
	"strconv"
	"sync"
)

type Config struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	From     string
	Subject  string
}

type Service struct {
	ratesSubscribersService ports.RatesSubscribersService
	exchangeRatesService    ports.ExchangeRatesService
	*Config
}

func NewService(
	config *Config,
	ratesSubscribersService ports.RatesSubscribersService,
	exchangeRatesService ports.ExchangeRatesService,
) *Service {
	return &Service{
		ratesSubscribersService: ratesSubscribersService,
		exchangeRatesService:    exchangeRatesService,
		Config:                  config,
	}
}

func (s *Service) SendRates(ctx context.Context) error {
	subscribers, err := s.ratesSubscribersService.GetAllSubscribers(ctx)
	if err != nil {
		return err
	}
	rate, err := s.exchangeRatesService.GetUAHUSDExchangeRate(ctx)
	if err != nil {
		return err
	}
	str := strconv.FormatFloat(rate, 'f', 2, 64)

	var wg sync.WaitGroup
	errCh := make(chan error, len(subscribers))

	for _, sub := range subscribers {
		wg.Add(1)
		go func(subscriber subscriber.Subscriber) {
			defer wg.Done()
			if err := s.send(ctx, str, subscriber); err != nil {
				errCh <- err
			}
		}(sub)
	}

	wg.Wait()
	close(errCh)

	if len(errCh) == 0 {
		return nil
	}
	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errNotification(errs)
	}

	return nil
}

// func (s *Service) send(ctx context.Context, message string, subscriber subscriber.Subscriber) error {
// 	to := []string{subscriber.Email}
//
// 	m := []byte(s.Subject + "\n" + message)
//
// 	auth := smtp.PlainAuth("", s.Username, s.Password, s.SMTPHost)
//
// 	conn, err := smtp.Dial(s.SMTPHost + ":" + s.SMTPPort)
// 	if err != nil {
// 		return errSentNotification(err)
// 	}
// 	defer conn.Close()
//
// 	if err := conn.StartTLS(nil); err != nil {
// 		return errSentNotification(err)
// 	}
//
// 	if err := conn.Auth(auth); err != nil {
// 		return errSentNotification(err)
// 	}
//
// 	if err := conn.Mail(s.From); err != nil {
// 		return errSentNotification(err)
// 	}
//
// 	for _, addr := range to {
// 		if err := conn.Rcpt(addr); err != nil {
// 			return errSentNotification(err)
// 		}
// 	}
//
// 	w, err := conn.Data()
// 	if err != nil {
// 		return errSentNotification(err)
// 	}
//
// 	_, err = w.Write(m)
// 	if err != nil {
// 		return errSentNotification(err)
// 	}
//
// 	err = w.Close()
// 	if err != nil {
// 		return errSentNotification(err)
// 	}
//
// 	return nil
// }

func (s *Service) send(ctx context.Context, message string, subscriber subscriber.Subscriber) error {
	to := []string{subscriber.Email}
	m := []byte(s.Subject + "\n" + message)

	auth := smtp.PlainAuth("", s.Username, s.Password, s.SMTPHost)
	err := smtp.SendMail(s.SMTPHost+":"+s.SMTPPort, auth, s.From, to, m)
	if err != nil {
		return err
	}
	return nil
}
