package customerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm/constants"
	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/FUJI0130/curriculum/src/core/support/errorcodes"
)

//career

type CareerError struct {
	base.BaseError
}

func ErrInvalidCareerDetail(cause error, customMsg ...string) error {
	fullMessage := "Invalid career detail provided"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &CareerError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

type CareerIDError struct {
	base.BaseError
}

func ErrInvalidCareerIDFormat(cause error, customMsg ...string) error {
	fullMessage := "Invalid career ID format provided"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &CareerIDError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

//Skill

type SkillError struct {
	base.BaseError
}

func ErrInvalidSkillEvaluation(cause error, value uint8, customMsg ...string) error {
	fullMessage := fmt.Sprintf("Invalid skill evaluation provided: %d", value)
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &SkillError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

func ErrInvalidSkillYear(cause error, year uint8, customMsg ...string) error {
	fullMessage := fmt.Sprintf("Invalid skill year provided: %d", year)
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &SkillError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

type SkillYearError struct {
	base.BaseError
}

func ErrSkillYearZeroOrNegative(cause error, customMsg ...string) error {
	fullMessage := "SkillYear cannot be zero or negative value"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &SkillYearError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

func ErrSkillYearTooLong(cause error, customMsg ...string) error {
	fullMessage := "SkillYear is too long"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &SkillYearError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

type SkillIDError struct {
	base.BaseError
}

func ErrInvalidSkillIDFormat(cause error, customMsg ...string) error {
	fullMessage := "Invalid skill ID format provided"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &SkillIDError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

type SkillEvaluationError struct {
	base.BaseError
}

func ErrSkillEvaluationOutOfRange(cause error, value uint8, customMsg ...string) error {
	fullMessage := fmt.Sprintf("SkillEvaluation must be between %d and %d", constants.MinSkillEvaluationValue, constants.MaxSkillEvaluationValue)
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &SkillEvaluationError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

// user
type UserIDError struct {
	base.BaseError
}

func ErrInvalidUserIDFormat(cause error, customMsg ...string) error {
	fullMessage := "Invalid user ID format provided"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserIDError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

type UserEmailError struct {
	base.BaseError
}

func ErrUserEmailEmpty(cause error, customMsg ...string) error {
	fullMessage := "UserEmail cannot be empty"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserEmailError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

func ErrUserEmailTooLong(cause error, customMsg ...string) error {
	fullMessage := "UserEmail length over maximum allowed length"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserEmailError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

func ErrUserEmailInvalidFormat(cause error, customMsg ...string) error {
	fullMessage := "UserEmail format is invalid"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserEmailError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

type UserPasswordError struct {
	base.BaseError
}

func ErrUserPasswordEmpty(cause error, customMsg ...string) error {
	fullMessage := "UserPassword cannot be empty"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserPasswordError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

func ErrUserPasswordTooLong(cause error, customMsg ...string) error {
	fullMessage := "UserPassword length over maximum allowed length"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserPasswordError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

func ErrUserPasswordTooShort(cause error, customMsg ...string) error {
	fullMessage := "UserPassword length under minimum required length"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserPasswordError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

type UserProfileError struct {
	base.BaseError
}

func ErrUserProfileEmpty(cause error, customMsg ...string) error {
	fullMessage := "User profile cannot be empty"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserProfileError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

func ErrUserProfileTooLong(cause error, customMsg ...string) error {
	fullMessage := "User profile length over maximum allowed length"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserProfileError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}

type UserNotFoundError struct {
	base.BaseError
}

func (e *UserNotFoundError) Unwrap() error {
	return &e.BaseError
}

func (e *UserNotFoundError) Is(target error) bool {
	_, ok := target.(*UserNotFoundError)
	return ok
}

func ErrUserNotFound(cause error, customMsg ...string) error {
	fullMessage := "User could not be found in the database"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserNotFoundError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.NotFound, cause),
	}
}

// Test function remains unchanged
func TestErrorIs(t *testing.T) {
	err := ErrUserNotFound(nil)
	if !errors.Is(err, ErrUserNotFound(nil)) {
		t.Fatal("errors.Is did not recognize ErrUserNotFound")
	}
}

type UserNameAlreadyExistsError struct {
	base.BaseError
}

func ErrUserNameAlreadyExists(cause error, name string, customMsg ...string) error {
	fullMessage := fmt.Sprintf("Name provided: %s", name)
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &UserNameAlreadyExistsError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}
