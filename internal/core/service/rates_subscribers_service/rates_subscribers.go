package rates_subscribers_service

import (
	"context"
	"github.com/sergey4qb/rate/internal/core/domain/subscriber"
	"github.com/sergey4qb/rate/internal/core/ports"
)

type Service struct {
	db ports.RatesSubscribersStorage
}

func NewService(db ports.RatesSubscribersStorage) *Service {
	return &Service{db: db}
}

func (s *Service) SaveSubscriber(ctx context.Context, subscriber subscriber.Subscriber) error {
	if err := subscriber.Validate(); err != nil {
		return err
	}
	exist, err := s.db.SubscriberExists(ctx, subscriber.Email)
	if err != nil {
		return err
	}
	if exist {
		return ErrSubscriberAlreadyExist
	}
	return s.db.SaveSubscriber(ctx, subscriber)
}

func (s *Service) GetAllSubscribers(ctx context.Context) ([]subscriber.Subscriber, error) {
	subscribers, err := s.db.GetSubscribers(ctx)
	if err != nil {
		return nil, err
	}
	if subscribers == nil {
		return nil, ErrSubscribersNotFound
	}
	return subscribers, nil
}
