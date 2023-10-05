package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm/constants"
	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/FUJI0130/curriculum/src/core/support/errorcodes"
)

type TagIDError struct {
	base.BaseError
}

// customMsgをスライスにしている部分を無くしたい
func ErrInvalidTagIDFormat(cause error, customMsg string) error {
	fullMessage := "Invalid tag ID format provided"
	if customMsg != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg)
	} else {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &TagIDError{
		BaseError: *base.NewBaseError(
			fullMessage,
			errorcodes.BadRequest,
			cause,
		),
	}
}

type TagNameError struct {
	base.BaseError
}

func ErrTagNameEmpty(cause error, customMsg ...string) error {
	fullMessage := "Tag name cannot be empty"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &TagNameError{
		BaseError: *base.NewBaseError(
			fullMessage,
			errorcodes.BadRequest,
			cause,
		),
	}
}

type TagNameTooLongError struct {
	base.BaseError
}

func ErrTagNameTooLong(cause error, customMsg ...string) error {
	fullMessage := fmt.Sprintf("Tag name must be less than or equal to %d characters", constants.NameMaxLength)
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &TagNameTooLongError{
		BaseError: *base.NewBaseError(
			fullMessage,
			errorcodes.BadRequest,
			cause,
		),
	}
}

type TagNameAlreadyExistsError struct {
	base.BaseError
}

func ErrTagNameAlreadyExists(cause error, name string, customMsg ...string) error {
	fullMessage := fmt.Sprintf("Tag name provided: %s", name)
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &TagNameAlreadyExistsError{
		BaseError: *base.NewBaseError(
			fullMessage,
			errorcodes.BadRequest,
			cause,
		),
	}
}

type DuplicateSkillTagError struct {
	base.BaseError
}

func ErrDuplicateSkillTag(cause error, name string, customMsg ...string) error {
	fullMessage := fmt.Sprintf("Duplicate skill tag provided: %s", name)
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &DuplicateSkillTagError{
		BaseError: *base.NewBaseError(
			fullMessage,
			errorcodes.BadRequest,
			cause,
		),
	}
}
