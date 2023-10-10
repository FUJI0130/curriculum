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
			TraceVal:      errors.New(message),
		},
	}
}

func NewInternalServerErrorf(format string, args ...any) *InternalServerErrorType {
	message := fmt.Sprintf(format, args...)
	return &InternalServerErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeInternalServerError,
			TraceVal:      errors.New(message),
		},
	}
}

func WrapInternalServerError(err error, message string) *InternalServerErrorType {
	combinedMessage := fmt.Sprintf("%s: %s", message, err.Error())
	return &InternalServerErrorType{
		&BaseErr{
			Message:       combinedMessage,
			StatusCodeVal: errCodeInternalServerError,
			TraceVal:      errors.Wrap(err, combinedMessage),
		},
	}
}

func WrapInternalServerErrorf(err error, format string, args ...any) *InternalServerErrorType {
	extraMessage := fmt.Sprintf(format, args...)
	combinedMessage := fmt.Sprintf("%s: %s", extraMessage, err.Error())
	return &InternalServerErrorType{
		&BaseErr{
			Message:       combinedMessage,
			StatusCodeVal: errCodeInternalServerError,
			TraceVal:      errors.Wrap(err, combinedMessage),
		},
	}
}
