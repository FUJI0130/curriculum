package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/common/base"
)

type UserNameAlreadyExistsError struct {
	base.BaseError
}

func ErrUserNameAlreadyExists(name string) error {
	return &UserNameAlreadyExistsError{
		BaseError: *base.NewBaseError(
			fmt.Sprintf("User name provided: %s", name),
			400,
			"User name already exists",
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
			400,
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
			400,
			"Cannot have duplicate skill tags",
		),
	}
}

type DatabaseError struct {
	base.BaseError
}

func ErrDatabaseError(detail string) error {
	return &DatabaseError{
		BaseError: *base.NewBaseError(
			detail,
			500,
			"Database operation error",
		),
	}
}
