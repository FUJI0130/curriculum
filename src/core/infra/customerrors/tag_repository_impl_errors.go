package customerrors

import (
	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
)

type TagNotFoundError struct {
	base.BaseError
}

func ErrTagNotFound() error {
	return &TagNotFoundError{
		BaseError: *base.NewBaseError(
			"Tag not found",
			errorcodes.NotFound,
			"The specified tag could not be found in the database",
		),
	}
}
