package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/cockroachdb/errors"
)

// UnprocessableEntityError
type UnprocessableEntityErrorType struct {
	*base.BaseError
}

func NewUnprocessableEntityError(message string) *UnprocessableEntityErrorType {
	return &UnprocessableEntityErrorType{
		&base.BaseError{
			Message:       message,
			StatusCodeVal: ErrCodeUnprocessableEntity,
			TraceVal:      errors.New(message),
		},
	}
}

func NewUnprocessableEntityErrorf(format string, args ...interface{}) *UnprocessableEntityErrorType {
	message := fmt.Sprintf(format, args...)
	return &UnprocessableEntityErrorType{
		&base.BaseError{
			Message:       message,
			StatusCodeVal: ErrCodeUnprocessableEntity,
			TraceVal:      errors.New(message),
		},
	}
}

func WrapUnprocessableEntityError(err error, message string) *UnprocessableEntityErrorType {
	return &UnprocessableEntityErrorType{
		&base.BaseError{
			Message:       message,
			StatusCodeVal: ErrCodeUnprocessableEntity,
			TraceVal:      errors.Wrap(err, message),
		},
	}
}
func WrapUnprocessableEntityErrorf(err error, format string, args ...interface{}) *UnprocessableEntityErrorType {
	message := fmt.Sprintf(format, args...)
	return &UnprocessableEntityErrorType{
		&base.BaseError{
			Message:       message,
			StatusCodeVal: ErrCodeUnprocessableEntity,
			TraceVal:      errors.Wrap(err, message),
		},
	}
}
