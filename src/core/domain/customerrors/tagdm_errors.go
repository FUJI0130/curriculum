package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
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
			errorcodes.BadRequest,
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
			errorcodes.BadRequest,
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
			errorcodes.BadRequest,
			fmt.Sprintf("Tag name must be less than or equal to %d characters", constants.NameMaxLength),
		),
	}
}

type TagNameAlreadyExistsError struct {
	base.BaseError
}

func ErrTagNameAlreadyExists(name string) error {
	return &TagNameAlreadyExistsError{
		BaseError: *base.NewBaseError(
			fmt.Sprintf("Tag name provided: %s", name),
			errorcodes.BadRequest,
			"Tag name already exists",
		),
	}
}

type DuplicateSkillTagError struct {
	base.BaseError
}

func ErrDuplicateSkillTag(name string) error {
	return &DuplicateSkillTagError{
		BaseError: *base.NewBaseError(
			fmt.Sprintf("Duplicate skill tag provided: %s", name),
			errorcodes.BadRequest,
			"Cannot have duplicate skill tags",
		),
	}
}
