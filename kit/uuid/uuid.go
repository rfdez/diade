package uuid

import "github.com/google/uuid"

// EnsureIsValidUUID ensures that the given string is a valid UUID.
func EnsureIsValidUUID(s string) error {
	_, err := uuid.Parse(s)
	if err != nil {
		return err
	}

	return nil
}
