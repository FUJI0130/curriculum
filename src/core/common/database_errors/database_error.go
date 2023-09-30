package database_errors

import (
	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
)

type DatabaseError struct {
	base.BaseError
}

func ErrDatabaseError(detail string) error {
	return &DatabaseError{
		BaseError: *base.NewBaseError(
			"Database error occurred",
			errorcodes.InternalServerError,
			detail,
		),
	}
}
