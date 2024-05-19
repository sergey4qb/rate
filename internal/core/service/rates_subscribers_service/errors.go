package rates_subscribers_service

import "errors"

var (
	ErrSubscriberAlreadyExist = errors.New("subscriber already exists")
	ErrSubscribersNotFound    = errors.New("subscribers not found")
)
