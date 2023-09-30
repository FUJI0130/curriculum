package customerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm/constants"
)

//career

type CareerError struct {
	base.BaseError
}

func ErrInvalidCareerDetail(detail string) error {
	return &CareerError{
		BaseError: *base.NewBaseError("Invalid career detail provided", errorcodes.BadRequest, "Detail provided: "+detail),
	}
}

type CareerIDError struct {
	base.BaseError
}

func ErrInvalidCareerIDFormat(id string) error {
	return &CareerIDError{
		BaseError: *base.NewBaseError("Invalid career ID format provided", errorcodes.BadRequest, "ID provided: "+id),
	}
}

//Skill

type SkillError struct {
	base.BaseError
}

func ErrInvalidSkillEvaluation(value uint8) error {
	return &SkillError{
		BaseError: *base.NewBaseError(
			"Invalid skill evaluation provided",
			errorcodes.BadRequest,
			fmt.Sprintf("Evaluation provided: %d", value),
		),
	}
}

func ErrInvalidSkillYear(year uint8) error {
	return &SkillError{
		BaseError: *base.NewBaseError(
			"Invalid skill year provided",
			errorcodes.BadRequest,
			fmt.Sprintf("Year provided: %d", year),
		),
	}
}

type SkillYearError struct {
	base.BaseError
}

func ErrSkillYearZeroOrNegative() error {
	return &SkillYearError{
		BaseError: *base.NewBaseError(
			"SkillYear cannot be zero or negative value",
			errorcodes.BadRequest,
			"SkillYear should be a positive value",
		),
	}
}

func ErrSkillYearTooLong() error {
	return &SkillYearError{
		BaseError: *base.NewBaseError(
			"SkillYear is too long",
			errorcodes.BadRequest,
			"SkillYear should be less than or equal to 100",
		),
	}
}

type SkillIDError struct {
	base.BaseError
}

func ErrInvalidSkillIDFormat() error {
	return &SkillIDError{
		BaseError: *base.NewBaseError(
			"Invalid skill ID format provided",
			errorcodes.BadRequest,
			"Invalid UUID format",
		),
	}
}

type SkillEvaluationError struct {
	base.BaseError
}

func ErrSkillEvaluationOutOfRange(value uint8) error {
	return &SkillEvaluationError{
		BaseError: *base.NewBaseError(
			fmt.Sprintf("SkillEvaluation must be between %d and %d", constants.MinSkillEvaluationValue, constants.MaxSkillEvaluationValue),
			errorcodes.BadRequest,
			fmt.Sprintf("Evaluation provided: %d", value),
		),
	}
}

// user
type UserIDError struct {
	base.BaseError
}

func ErrInvalidUserIDFormat() error {
	return &UserIDError{
		BaseError: *base.NewBaseError(
			"Invalid user ID format provided",
			errorcodes.BadRequest,
			"Invalid UUID format",
		),
	}
}

type UserEmailError struct {
	base.BaseError
}

func ErrUserEmailEmpty() error {
	return &UserEmailError{
		BaseError: *base.NewBaseError(
			"UserEmail cannot be empty",
			errorcodes.BadRequest,
			"Provided email is empty",
		),
	}
}

func ErrUserEmailTooLong() error {
	return &UserEmailError{
		BaseError: *base.NewBaseError(
			"UserEmail length over maximum allowed length",
			errorcodes.BadRequest,
			"Provided email exceeds the maximum length",
		),
	}
}

func ErrUserEmailInvalidFormat() error {
	return &UserEmailError{
		BaseError: *base.NewBaseError(
			"UserEmail format is invalid",
			errorcodes.BadRequest,
			"Invalid email format",
		),
	}
}

type UserPasswordError struct {
	base.BaseError
}

func ErrUserPasswordEmpty() error {
	return &UserPasswordError{
		BaseError: *base.NewBaseError(
			"UserPassword cannot be empty",
			errorcodes.BadRequest,
			"Provided password is empty",
		),
	}
}

func ErrUserPasswordTooLong() error {
	return &UserPasswordError{
		BaseError: *base.NewBaseError(
			"UserPassword length over maximum allowed length",
			errorcodes.BadRequest,
			"Provided password exceeds the maximum length",
		),
	}
}

func ErrUserPasswordTooShort() error {
	return &UserPasswordError{
		BaseError: *base.NewBaseError(
			"UserPassword length under minimum required length",
			errorcodes.BadRequest,
			"Provided password is too short",
		),
	}
}

type UserProfileError struct {
	base.BaseError
}

func ErrUserProfileEmpty() error {
	return &UserProfileError{
		BaseError: *base.NewBaseError(
			"User profile cannot be empty",
			errorcodes.BadRequest,
			"Provided profile is empty",
		),
	}
}

func ErrUserProfileTooLong() error {
	return &UserProfileError{
		BaseError: *base.NewBaseError(
			"User profile length over maximum allowed length",
			errorcodes.BadRequest,
			"Provided profile exceeds the maximum length",
		),
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

func ErrUserNotFound() error {
	return &UserNotFoundError{
		BaseError: *base.NewBaseError(
			"User could not be found in the database",
			errorcodes.NotFound,
			"User not found",
		),
	}
}

func TestErrorIs(t *testing.T) {
	err := ErrUserNotFound()
	if !errors.Is(err, ErrUserNotFound()) {
		t.Fatal("errors.Is did not recognize ErrUserNotFound")
	}
}

type UserNameAlreadyExistsError struct {
	base.BaseError
}

func ErrUserNameAlreadyExists(name string) error {
	return &UserNameAlreadyExistsError{
		BaseError: *base.NewBaseError(
			fmt.Sprintf("Name provided: %s", name),
			errorcodes.BadRequest,
			"User name already exists",
		),
	}
}
