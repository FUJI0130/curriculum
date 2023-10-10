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
			Message:       message,
			StatusCodeVal: ErrCodeNotFound,
			TraceVal:      errors.New(message),
		},
	}
}

func NewNotFoundErrorf(format string, args ...interface{}) *NotFoundErrorType {
	message := fmt.Sprintf(format, args...)
	return &NotFoundErrorType{
		&base.BaseError{
			Message:       message,
			StatusCodeVal: ErrCodeNotFound,
			TraceVal:      errors.New(message),
		},
	}
}

func WrapNotFoundError(err error, message string) *NotFoundErrorType {
	return &NotFoundErrorType{
		&base.BaseError{
			Message:       message,
			StatusCodeVal: ErrCodeNotFound,
			TraceVal:      errors.Wrap(err, message),
		},
	}
}
func WrapNotFoundErrorf(err error, format string, args ...interface{}) *NotFoundErrorType {
	message := fmt.Sprintf(format, args...)
	return &NotFoundErrorType{
		&base.BaseError{
			Message:       message,
			StatusCodeVal: ErrCodeNotFound,
			TraceVal:      errors.Wrap(err, message),
		},
	}
}
