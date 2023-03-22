package datetime

import (
	"time"
)

// EnsureIsValidDate ensures that the given string is a valid date.
func EnsureIsValidDate(s string) error {
	_, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	return nil
}
