package ports

import "context"

type RateNotificationService interface {
	SendRates(ctx context.Context) error
}
