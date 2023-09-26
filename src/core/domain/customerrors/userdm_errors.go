package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm/constants"
)

//career

type CareerError struct {
	base.BaseError
}

func ErrInvalidCareerDetail(detail string) error {
	return &CareerError{
		BaseError: *base.NewBaseError("Invalid career detail provided", 400, "Detail provided: "+detail),
	}
}

type CareerIDError struct {
	base.BaseError
}

func ErrInvalidCareerIDFormat(id string) error {
	return &CareerIDError{
		BaseError: *base.NewBaseError("Invalid career ID format provided", 400, "ID provided: "+id),
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
			400,
			fmt.Sprintf("Evaluation provided: %d", value),
		),
	}
}

func ErrInvalidSkillYear(year uint8) error {
	return &SkillError{
		BaseError: *base.NewBaseError(
			"Invalid skill year provided",
			400,
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
			400,
			"SkillYear should be a positive value",
		),
	}
}

func ErrSkillYearTooLong() error {
	return &SkillYearError{
		BaseError: *base.NewBaseError(
			"SkillYear is too long",
			400,
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
			400,
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
			400,
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
			400,
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
			400,
			"Provided email is empty",
		),
	}
}

func ErrUserEmailTooLong() error {
	return &UserEmailError{
		BaseError: *base.NewBaseError(
			"UserEmail length over maximum allowed length",
			400,
			"Provided email exceeds the maximum length",
		),
	}
}

func ErrUserEmailInvalidFormat() error {
	return &UserEmailError{
		BaseError: *base.NewBaseError(
			"UserEmail format is invalid",
			400,
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
			400,
			"Provided password is empty",
		),
	}
}

func ErrUserPasswordTooLong() error {
	return &UserPasswordError{
		BaseError: *base.NewBaseError(
			"UserPassword length over maximum allowed length",
			400,
			"Provided password exceeds the maximum length",
		),
	}
}

func ErrUserPasswordTooShort() error {
	return &UserPasswordError{
		BaseError: *base.NewBaseError(
			"UserPassword length under minimum required length",
			400,
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
			400,
			"Provided profile is empty",
		),
	}
}

func ErrUserProfileTooLong() error {
	return &UserProfileError{
		BaseError: *base.NewBaseError(
			"User profile length over maximum allowed length",
			400,
			"Provided profile exceeds the maximum length",
		),
	}
}

type UserNotFoundError struct {
	base.BaseError
}

func ErrUserNotFound() error {
	return &UserNotFoundError{
		BaseError: *base.NewBaseError(
			"User could not be found in the database",
			404,
			"User not found",
		),
	}
}

type UserNameAlreadyExistsError struct {
	base.BaseError
}

func ErrUserNameAlreadyExists(name string) error {
	return &UserNameAlreadyExistsError{
		BaseError: *base.NewBaseError(
			fmt.Sprintf("Name provided: %s", name),
			400,
			"User name already exists",
		),
	}
}
