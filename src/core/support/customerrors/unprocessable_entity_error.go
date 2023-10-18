package customerrors

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

const errCodeUnprocessableEntity = 400

type UnprocessableEntityErrorType struct {
	*BaseErr
}

func NewUnprocessableEntityError(message string) error {
	return &UnprocessableEntityErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeUnprocessableEntity,
			TraceVal:      errors.Errorf("%+v", errors.New(message)),
		},
	}
}

func NewUnprocessableEntityErrorf(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	return &UnprocessableEntityErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeUnprocessableEntity,
			TraceVal:      errors.Errorf("%+v", errors.New(message)),
		},
	}
}

func WrapUnprocessableEntityError(err error, message string) error {
	baseError := NewBaseError(message, errCodeUnprocessableEntity, nil)
	wrappedError := baseError.WrapWithLocation(err, message)
	return &UnprocessableEntityErrorType{
		BaseErr: wrappedError,
	}
}

func WrapUnprocessableEntityErrorf(err error, format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	baseError := NewBaseError(message, errCodeUnprocessableEntity, nil)
	wrappedError := baseError.WrapWithLocation(err, message)
	return &UnprocessableEntityErrorType{
		BaseErr: wrappedError,
	}
}
