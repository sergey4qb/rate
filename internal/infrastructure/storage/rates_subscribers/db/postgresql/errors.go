package postgresql

import "fmt"

func errFailToInsertSubscriber(reason error) error {
	return fmt.Errorf("db: failed to insert subscriber | reason: %w", reason)
}

func errGetSubscribers(reason error) error {
	return fmt.Errorf("db: failed to retrieve subscribers | reason: %w", reason)
}

func errSubscriberExist(reason error) error {
	return fmt.Errorf("db: failed to check if subscriber exists | reason: %w", reason)
}
