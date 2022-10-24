package datetime

import (
	"time"
)

func EnsureIsValidDate(s string) error {
	_, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	return nil
}
