package ports

import (
	"context"
	"github.com/sergey4qb/rate/internal/core/domain/subscriber"
)

type RatesSubscribersStorage interface {
	SaveSubscriber(ctx context.Context, subscriber subscriber.Subscriber) error
	GetSubscribers(ctx context.Context) ([]subscriber.Subscriber, error)
	SubscriberExists(ctx context.Context, email string) (bool, error)
}

type RatesSubscribersService interface {
	SaveSubscriber(ctx context.Context, subscriber subscriber.Subscriber) error
	GetAllSubscribers(ctx context.Context) ([]subscriber.Subscriber, error)
}
