package uuid

import "github.com/google/uuid"

func EnsureIsValidUUID(s string) error {
	_, err := uuid.Parse(s)
	if err != nil {
		return err
	}

	return nil
}
