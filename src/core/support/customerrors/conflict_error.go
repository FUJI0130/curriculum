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
			TraceVal:      errors.WithStack(errors.New(message)),
		},
	}
}

func NewConflictErrorf(format string, args ...any) *ConflictErrorType {
	message := fmt.Sprintf(format, args...)
	return &ConflictErrorType{
		&BaseErr{
			Message:       message,
			StatusCodeVal: errCodeConflict,
			TraceVal:      errors.WithStack(errors.New(message)),
		},
	}
}

func WrapConflictError(err error, message string) *ConflictErrorType {
	baseError := NewBaseError(message, errCodeConflict, nil)
	wrappedError := baseError.WrapWithLocation(err, message)
	return &ConflictErrorType{
		BaseErr: wrappedError,
	}
}

func WrapConflictErrorf(err error, format string, args ...interface{}) *ConflictErrorType {
	message := fmt.Sprintf(format, args...)
	baseError := NewBaseError(message, errCodeConflict, nil)
	wrappedError := baseError.WrapWithLocation(err, message)
	return &ConflictErrorType{
		BaseErr: wrappedError,
	}
}
