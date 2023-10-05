package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/cockroachdb/errors"
)

// BadRequestError
type BadRequestErrorType struct {
	*base.BaseError
}

func NewBadRequestError(message string) *BadRequestErrorType {
	return &BadRequestErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeBadRequest,
			Trace:      errors.New(message),
		},
	}
}

func NewBadRequestErrorf(format string, args ...interface{}) *BadRequestErrorType {
	message := fmt.Sprintf(format, args...)
	return &BadRequestErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeBadRequest,
			Trace:      errors.New(message),
		},
	}
}

func WrapBadRequestError(err error, message string) *BadRequestErrorType {
	return &BadRequestErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeBadRequest,
			Trace:      errors.Wrap(err, message),
		},
	}
}
