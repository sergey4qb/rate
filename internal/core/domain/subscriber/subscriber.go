package subscriber

import (
	"regexp"
)

type Subscriber struct {
	Email string
}

func (s *Subscriber) Validate() error {
	if err := validateEmail(s.Email); err != nil {
		return err
	}
	return nil
}

func validateEmail(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}
