package postgresql

import (
	"context"
	"github.com/sergey4qb/rate/internal/core/domain/subscriber"
)

func (p *PostgreQuery) SaveSubscriber(ctx context.Context, subscriber subscriber.Subscriber) error {
	query := `
        INSERT INTO rates_subscribers (email) 
        VALUES ($1)
        ON CONFLICT (email) DO NOTHING
    `
	if _, err := p.conn.ExecContext(ctx, query, subscriber.Email); err != nil {
		return errFailToInsertSubscriber(err)
	}
	return nil
}

func (p *PostgreQuery) GetSubscribers(ctx context.Context) ([]subscriber.Subscriber, error) {
	query := `SELECT email FROM rates_subscribers`

	rows, err := p.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, errGetSubscribers(err)
	}
	defer rows.Close()

	var subscribers []subscriber.Subscriber
	for rows.Next() {
		var subscriber subscriber.Subscriber
		if err := rows.Scan(&subscriber.Email); err != nil {
			return nil, errGetSubscribers(err)
		}
		subscribers = append(subscribers, subscriber)
	}
	if err := rows.Err(); err != nil {
		return nil, errGetSubscribers(err)
	}
	return subscribers, nil
}

func (p *PostgreQuery) SubscriberExists(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM rates_subscribers WHERE email = $1)`

	var exists bool
	err := p.conn.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, errSubscriberExist(err)
	}
	return exists, nil
}
