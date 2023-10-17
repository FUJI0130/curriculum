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
	// _, file, line, _ := runtime.Caller(1)
	// formattedMessage := fmt.Sprintf("[File: %s Line: %d] %s", file, line, message)

	return &UnprocessableEntityErrorType{
		&BaseErr{
			// Message:       formattedMessage,
			Message:       message,
			StatusCodeVal: errCodeUnprocessableEntity,
			// TraceVal:      errors.WithStack(errors.New(formattedMessage)),
			TraceVal: errors.WithStack(errors.New(message)),
			// TraceVal: errors.WithStack(fmt.Errorf("%s", message)),
			// TraceVal: errors.WithStackDepth(errors.New(message), 0),
			// TraceVal: errors.New(message),
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
