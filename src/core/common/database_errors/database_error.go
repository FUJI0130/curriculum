package database_errors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
)

type DatabaseError struct {
	base.BaseError
}

func ErrDatabaseError(cause error, customMsg ...string) *base.BaseError {
	fullMessage := "Database error occurred"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return base.NewBaseError(
		fullMessage,
		errorcodes.InternalServerError,
		cause,
	)
}
