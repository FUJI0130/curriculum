package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/cockroachdb/errors"
)

// ConflictError
type ConflictErrorType struct {
	*base.BaseError
}

func NewConflictError(message string) *ConflictErrorType {
	return &ConflictErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeConflict,
			Trace:      errors.New(message),
		},
	}
}

func NewConflictErrorf(format string, args ...interface{}) *ConflictErrorType {
	message := fmt.Sprintf(format, args...)
	return &ConflictErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeConflict,
			Trace:      errors.New(message),
		},
	}
}

func WrapConflictError(err error, message string) *ConflictErrorType {
	return &ConflictErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeConflict,
			Trace:      errors.Wrap(err, message),
		},
	}
}

func WrapConflictErrorf(err error, format string, args ...interface{}) *ConflictErrorType {
	message := fmt.Sprintf(format, args...)
	return &ConflictErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeConflict,
			Trace:      errors.Wrapf(err, format, args...),
		},
	}
}
