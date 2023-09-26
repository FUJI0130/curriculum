package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm/constants"
)

// Tag errors
type TagIDError struct {
	base.BaseError
}

func ErrInvalidTagIDFormat() error {
	return &TagIDError{
		BaseError: *base.NewBaseError(
			"Invalid tag ID format provided",
			400,
			"Invalid UUID format",
		),
	}
}

type TagNameError struct {
	base.BaseError
}

func ErrTagNameEmpty() error {
	return &TagNameError{
		BaseError: *base.NewBaseError(
			"Tag name cannot be empty",
			400,
			"Provided tag name is empty",
		),
	}
}

type TagNameTooLongError struct {
	base.BaseError
}

func ErrTagNameTooLong() error {
	return &TagNameTooLongError{
		BaseError: *base.NewBaseError(
			"Tag name is too long",
			400,
			fmt.Sprintf("Tag name must be less than or equal to %d characters", constants.NameMaxLength),
		),
	}
}
