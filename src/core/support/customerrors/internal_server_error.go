package customerrors

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

const errCodeInternalServerError = 500

// InternalServerError
type InternalServerErrorType struct {
	*BaseErr
}

func NewInternalServerError(message string) *InternalServerErrorType {
	return &InternalServerErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeInternalServerError,
			TraceVal:      errors.WithStack(errors.New(message)),
		},
	}
}

func NewInternalServerErrorf(format string, args ...any) *InternalServerErrorType {
	message := fmt.Sprintf(format, args...)
	return &InternalServerErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeInternalServerError,
			TraceVal:      errors.WithStack(errors.New(message)),
		},
	}
}

func WrapInternalServerError(err error, message string) *InternalServerErrorType {
	baseError := NewBaseError(message, errCodeInternalServerError, nil)
	wrappedError := baseError.WrapWithLocation(err, message)
	return &InternalServerErrorType{
		BaseErr: wrappedError,
	}
}

func WrapInternalServerErrorf(err error, format string, args ...any) *InternalServerErrorType {
	message := fmt.Sprintf(format, args...)
	baseError := NewBaseError(message, errCodeInternalServerError, nil)
	wrappedError := baseError.WrapWithLocation(err, message)
	return &InternalServerErrorType{
		BaseErr: wrappedError,
	}
}
