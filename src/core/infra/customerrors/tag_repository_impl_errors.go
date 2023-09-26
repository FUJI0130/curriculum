package customerrors

import (
	"github.com/FUJI0130/curriculum/src/core/common/base"
)

type TagNotFoundError struct {
	base.BaseError
}

func ErrTagNotFound() error {
	return &TagNotFoundError{
		BaseError: *base.NewBaseError(
			"Tag not found",
			404,
			"The specified tag could not be found in the database",
		),
	}
}
