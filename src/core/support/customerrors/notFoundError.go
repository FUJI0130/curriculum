package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/cockroachdb/errors"
)

// NotFoundError
type NotFoundErrorType struct {
	*base.BaseError
}

func NewNotFoundError(message string) *NotFoundErrorType {
	return &NotFoundErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeNotFound,
			Trace:      errors.New(message),
		},
	}
}

func NewNotFoundErrorf(format string, args ...interface{}) *NotFoundErrorType {
	message := fmt.Sprintf(format, args...)
	return &NotFoundErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeNotFound,
			Trace:      errors.New(message),
		},
	}
}

func WrapNotFoundError(err error, message string) *NotFoundErrorType {
	return &NotFoundErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeNotFound,
			Trace:      errors.Wrap(err, message),
		},
	}
}
