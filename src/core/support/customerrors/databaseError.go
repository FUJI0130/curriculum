package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/pkg/errors"
)

// ... 既存のappErrと関連エラータイプの定義 ...

type DatabaseErrorType struct {
	*base.BaseError
}

func (e *DatabaseErrorType) Error() string {
	return e.Message
}

func NewDatabaseError(message string) *DatabaseErrorType {
	return &DatabaseErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeDatabaseError,
			Trace:      errors.New(message),
		},
	}
}

func NewDatabaseErrorf(format string, args ...interface{}) *DatabaseErrorType {
	message := fmt.Sprintf(format, args...)
	return &DatabaseErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeDatabaseError,
			Trace:      errors.Errorf(format, args...),
		},
	}
}

func WrapDatabaseError(err error, format string, args ...interface{}) *DatabaseErrorType {
	message := fmt.Sprintf(format, args...)
	return &DatabaseErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeDatabaseError,
			Trace:      errors.Wrapf(err, format, args...),
		},
	}
}
func WrapDatabaseErrorf(err error, format string, args ...interface{}) *DatabaseErrorType {
	message := fmt.Sprintf(format, args...)
	return &DatabaseErrorType{
		&base.BaseError{
			Message:    message,
			StatusCode: ErrCodeDatabaseError,
			Trace:      errors.Wrapf(err, format, args...),
		},
	}
}
