package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/cockroachdb/errors"
)

// InternalServerError
type InternalServerErrorType struct {
	*base.BaseError
}

func NewInternalServerError(message string) *InternalServerErrorType {
	return &InternalServerErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeInternalServerError,
			Trace:      errors.New(message),
		},
	}
}

func NewInternalServerErrorf(format string, args ...interface{}) *InternalServerErrorType {
	message := fmt.Sprintf(format, args...)
	return &InternalServerErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeInternalServerError,
			Trace:      errors.New(message),
		},
	}
}

func WrapInternalServerError(err error, message string) *InternalServerErrorType {
	return &InternalServerErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeInternalServerError,
			Trace:      errors.Wrap(err, message),
		},
	}
}
func WrapInternalServerErrorf(err error, format string, args ...interface{}) *InternalServerErrorType {
	message := fmt.Sprintf(format, args...)
	return &InternalServerErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeInternalServerError,
			Trace:      errors.Wrap(err, message),
		},
	}
}
