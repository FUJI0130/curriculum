package customerrors

import (
	"github.com/FUJI0130/curriculum/src/core/common/base"
)

//createdAt

type CreatedAtError struct {
	base.BaseError
}

func ErrCreatedAtZeroValue() error {
	return &CreatedAtError{
		BaseError: *base.NewBaseError(
			"CreatedAt cannot be zero value",
			400,
			"Provided CreatedAt value is zero",
		),
	}
}

func ErrCreatedAtFutureValue() error {
	return &CreatedAtError{
		BaseError: *base.NewBaseError(
			"CreatedAt cannot be future date",
			400,
			"Provided CreatedAt value is in the future",
		),
	}
}

//updatedAt

type UpdatedAtError struct {
	base.BaseError
}

func ErrUpdatedAtZeroValue() error {
	return &UpdatedAtError{
		BaseError: *base.NewBaseError(
			"UpdatedAt cannot be zero value",
			400,
			"Provided UpdatedAt value is zero",
		),
	}
}
