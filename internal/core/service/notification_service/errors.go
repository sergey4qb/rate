package notification_service

import "fmt"

func errNotification(reason []error) error {
	return fmt.Errorf("failed sent notification | reason: %v", reason)
}
func errSentNotification(reason error) error {
	return fmt.Errorf("failed sent notification | reason: %v", reason)
}
