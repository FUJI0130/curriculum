package customerrors

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

const errCodeUnprocessableEntity = 400

type UnprocessableEntityErrorType struct {
	*BaseErr
}

func NewUnprocessableEntityError(message string) *UnprocessableEntityErrorType {
	return &UnprocessableEntityErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeUnprocessableEntity,
			TraceVal:      errors.WithStack(errors.New("")),
		},
	}
}

func NewUnprocessableEntityErrorf(format string, args ...any) *UnprocessableEntityErrorType {
	message := fmt.Sprintf(format, args...)
	return &UnprocessableEntityErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeUnprocessableEntity,
			TraceVal:      errors.WithStack(errors.New(message)),
		},
	}
}

func WrapUnprocessableEntityError(err error, message string) *UnprocessableEntityErrorType {
	baseError := NewBaseError(message, errCodeUnprocessableEntity, nil)
	wrappedError := baseError.WrapWithLocation(err, message)
	return &UnprocessableEntityErrorType{
		BaseErr: wrappedError,
	}
}

func WrapUnprocessableEntityErrorf(err error, format string, args ...interface{}) *UnprocessableEntityErrorType {
	message := fmt.Sprintf(format, args...)
	baseError := NewBaseError(message, errCodeUnprocessableEntity, nil)
	wrappedError := baseError.WrapWithLocation(err, message)
	return &UnprocessableEntityErrorType{
		BaseErr: wrappedError,
	}
}