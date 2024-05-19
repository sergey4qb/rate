package db

import (
	"context"
	"github.com/sergey4qb/rate/internal/core/domain/subscriber"
	"github.com/sergey4qb/rate/internal/infrastructure/storage/rates_subscribers/db/postgresql"
	"github.com/sergey4qb/rate/pkg/postgresql_database"
)

type Db struct {
	db *postgresql.PostgreQuery
}

func New(db *postgresql_database.DB) *Db {
	return &Db{db: postgresql.New(db)}
}

func (receiver *Db) SaveSubscriber(ctx context.Context, subscriber subscriber.Subscriber) error {
	return receiver.db.SaveSubscriber(ctx, subscriber)
}
func (receiver *Db) GetSubscribers(ctx context.Context) ([]subscriber.Subscriber, error) {
	return receiver.db.GetSubscribers(ctx)
}
func (receiver *Db) SubscriberExists(ctx context.Context, email string) (bool, error) {
	return receiver.db.SubscriberExists(ctx, email)
}
