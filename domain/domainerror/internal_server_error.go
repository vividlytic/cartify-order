package domainerror

import "github.com/pkg/errors"

type InternalServerError struct {
	message string
	cause   error
}

func NewInternalServerError(message string, cause error) error {
	return errors.WithStack(&InternalServerError{message, cause})
}

func (i *InternalServerError) Error() string {
	return i.message
}

func (i *InternalServerError) Unwrap() error {
	return i.cause
}
