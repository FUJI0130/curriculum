package customerrors

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

const errCodeNotFound = 404

type NotFoundErrorType struct {
	*BaseErr
}

func NewNotFoundError(message string) *NotFoundErrorType {
	return &NotFoundErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeNotFound,
			TraceVal:      errors.Errorf("%+v", errors.New(message)),
		},
	}
}

func NewNotFoundErrorf(format string, args ...any) *NotFoundErrorType {
	message := fmt.Sprintf(format, args...)
	return &NotFoundErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeNotFound,
			TraceVal:      errors.Errorf("%+v", errors.New(message)),
		},
	}
}

func WrapNotFoundError(err error, message string) *NotFoundErrorType {
	baseError := NewBaseError(message, errCodeNotFound, nil)
	wrappedError := baseError.WrapWithLocation(err, message)
	return &NotFoundErrorType{
		BaseErr: wrappedError,
	}
}

func WrapNotFoundErrorf(err error, format string, args ...interface{}) *NotFoundErrorType {
	message := fmt.Sprintf(format, args...)
	baseError := NewBaseError(message, errCodeNotFound, nil)
	wrappedError := baseError.WrapWithLocation(err, message)
	return &NotFoundErrorType{
		BaseErr: wrappedError,
	}
}

func IsNotFoundError(err error) bool {
	_, ok := err.(*NotFoundErrorType)
	return ok
}
