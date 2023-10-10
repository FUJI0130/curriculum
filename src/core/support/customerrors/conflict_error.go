package customerrors

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

// ConflictError
type ConflictErrorType struct {
	*BaseErr
}

const errCodeConflict = 409

func NewConflictError(message string) *ConflictErrorType {
	return &ConflictErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeConflict,
			TraceVal:      errors.New(message),
		},
	}
}

func NewConflictErrorf(format string, args ...any) *ConflictErrorType {
	message := fmt.Sprintf(format, args...)
	return &ConflictErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeConflict,
			TraceVal:      errors.New(message),
		},
	}
}

func WrapConflictError(err error, message string) *ConflictErrorType {
	combinedMessage := fmt.Sprintf("%s: %s", message, err.Error())
	return &ConflictErrorType{
		&BaseErr{
			Message:       combinedMessage,
			StatusCodeVal: errCodeConflict,
			TraceVal:      errors.Wrap(err, combinedMessage),
		},
	}
}

func WrapConflictErrorf(err error, format string, args ...any) *ConflictErrorType {
	extraMessage := fmt.Sprintf(format, args...)
	combinedMessage := fmt.Sprintf("%s: %s", extraMessage, err.Error())
	return &ConflictErrorType{
		&BaseErr{
			Message:       combinedMessage,
			StatusCodeVal: errCodeConflict,
			TraceVal:      errors.Wrap(err, combinedMessage),
		},
	}
}
