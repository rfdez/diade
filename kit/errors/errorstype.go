package errors

import "github.com/pkg/errors"

type notFound struct {
	error
}

// NewNotFound returns an error which wraps err that satisfies IsNotFound().
func NewNotFound(format string, args ...interface{}) error {
	return &notFound{errors.Errorf(format, args...)}
}

// WrapNotFound returns an error which wraps err that satisfies IsNotFound().
func WrapNotFound(err error, format string, args ...interface{}) error {
	return &notFound{errors.Wrapf(err, format, args...)}
}

// IsNotFound reports whether err was created with NewNotFound or WrapNotFound.
func IsNotFound(err error) bool {
	var target *notFound
	return errors.As(err, &target)
}

type wrongInput struct {
	error
}

// NewWrongInput returns an error which wraps err that satisfies IsWrongInput().
func NewWrongInput(format string, args ...interface{}) error {
	return &wrongInput{errors.Errorf(format, args...)}
}

// WrapWrongInput returns an error which wraps err that satisfies IsWrongInput().
func WrapWrongInput(err error, format string, args ...interface{}) error {
	return &wrongInput{errors.Wrapf(err, format, args...)}
}

// IsWrongInput reports whether err was created with NewWrongInput or WrapWrongInput.
func IsWrongInput(err error) bool {
	var target *wrongInput
	return errors.As(err, &target)
}

type notAuthorized struct {
	error
}

// NewNotAuthorized returns an error which wraps err that satisfies IsNotAuthorized().
func NewNotAuthorized(format string, args ...interface{}) error {
	return &notAuthorized{errors.Errorf(format, args...)}
}

// WrapNotAuthorized returns an error which wraps err that satisfies IsNotAuthorized().
func WrapNotAuthorized(err error, format string, args ...interface{}) error {
	return &notAuthorized{errors.Wrapf(err, format, args...)}
}

// IsNotAuthorized reports whether err was created with NewNotAuthorized or WrapNotAuthorized.
func IsNotAuthorized(err error) bool {
	var target *notAuthorized
	return errors.As(err, &target)
}
