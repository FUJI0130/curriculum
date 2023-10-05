package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/FUJI0130/curriculum/src/core/support/errorcodes"
)

//createdAt

type CreatedAtError struct {
	base.BaseError
}

func ErrCreatedAtZeroValue(cause error, customMsg ...string) error {
	defaultMessage := "CreatedAt cannot be zero value"
	detail := "Provided CreatedAt value is zero"
	if len(customMsg) > 0 && customMsg[0] != "" {
		detail = customMsg[0]
	} else if cause != nil {
		detail = fmt.Sprintf("%s: %v", detail, cause)
	}
	return &CreatedAtError{
		BaseError: *base.NewBaseError(defaultMessage, errorcodes.BadRequest, cause),
	}
}

func ErrCreatedAtFutureValue(cause error, customMsg ...string) error {
	defaultMessage := "CreatedAt cannot be future date"
	detail := "Provided CreatedAt value is in the future"
	if len(customMsg) > 0 && customMsg[0] != "" {
		detail = customMsg[0]
	} else if cause != nil {
		detail = fmt.Sprintf("%s: %v", detail, cause)
	}
	return &CreatedAtError{
		BaseError: *base.NewBaseError(defaultMessage, errorcodes.BadRequest, cause),
	}
}

// updatedAt

type UpdatedAtError struct {
	base.BaseError
}

func ErrUpdatedAtZeroValue(cause error, customMsg ...string) error {
	defaultMessage := "UpdatedAt cannot be zero value"
	detail := "Provided UpdatedAt value is zero"
	if len(customMsg) > 0 && customMsg[0] != "" {
		detail = customMsg[0]
	} else if cause != nil {
		detail = fmt.Sprintf("%s: %v", detail, cause)
	}
	return &UpdatedAtError{
		BaseError: *base.NewBaseError(defaultMessage, errorcodes.BadRequest, cause),
	}
}
