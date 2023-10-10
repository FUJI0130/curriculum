package customerrors

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

const errCodeNotFound = 404

// NotFoundError
type NotFoundErrorType struct {
	*BaseErr
}

func NewNotFoundError(message string) *NotFoundErrorType {
	return &NotFoundErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeNotFound,
			TraceVal:      errors.New(message),
		},
	}
}

func NewNotFoundErrorf(format string, args ...any) *NotFoundErrorType {
	message := fmt.Sprintf(format, args...)
	return &NotFoundErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeNotFound,
			TraceVal:      errors.New(message),
		},
	}
}

func WrapNotFoundError(err error, message string) *NotFoundErrorType {
	return &NotFoundErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeNotFound,
			TraceVal:      errors.Wrap(err, message),
		},
	}
}
func WrapNotFoundErrorf(err error, format string, args ...any) *NotFoundErrorType {
	message := fmt.Sprintf(format, args...)
	return &NotFoundErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeNotFound,
			TraceVal:      errors.Wrap(err, message),
		},
	}
}
