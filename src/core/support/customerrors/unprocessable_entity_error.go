package customerrors

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

const errCodeUnprocessableEntity = 400

// UnprocessableEntityError
type UnprocessableEntityErrorType struct {
	*BaseErr
}

func NewUnprocessableEntityError(message string) *UnprocessableEntityErrorType {
	return &UnprocessableEntityErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeUnprocessableEntity,
			TraceVal:      errors.WithStack(errors.New(message)),
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
	combinedMessage := fmt.Sprintf("%s: %s", message, err.Error())
	return &UnprocessableEntityErrorType{
		&BaseErr{
			Message:       combinedMessage,
			StatusCodeVal: errCodeUnprocessableEntity,
			TraceVal:      errors.Wrap(err, combinedMessage),
		},
	}
}

func WrapUnprocessableEntityErrorf(err error, format string, args ...any) *UnprocessableEntityErrorType {
	extraMessage := fmt.Sprintf(format, args...)
	combinedMessage := fmt.Sprintf("%s: %s", extraMessage, err.Error())
	return &UnprocessableEntityErrorType{
		&BaseErr{
			Message:       combinedMessage,
			StatusCodeVal: errCodeUnprocessableEntity,
			TraceVal:      errors.Wrap(err, combinedMessage),
		},
	}
}
