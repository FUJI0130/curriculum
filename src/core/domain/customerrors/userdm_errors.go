package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/common/base"
)

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

// customerrors/userdm_errors.go

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

// customerrors/userdm_errors.go

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

// customerrors/userdm_errors.go

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

// customerrors/userdm_errors.go
const (
	MinSkillEvaluationValue = 1
	MaxSkillEvaluationValue = 5
)

type SkillEvaluationError struct {
	base.BaseError
}

func ErrSkillEvaluationOutOfRange(value uint8) error {
	return &SkillEvaluationError{
		BaseError: *base.NewBaseError(
			fmt.Sprintf("SkillEvaluation must be between %d and %d", MinSkillEvaluationValue, MaxSkillEvaluationValue),
			400,
			fmt.Sprintf("Evaluation provided: %d", value),
		),
	}
}
