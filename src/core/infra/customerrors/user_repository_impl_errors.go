// customerrors/userdm_errors.go
package customerrors

import (
	"github.com/FUJI0130/curriculum/src/core/common/base"
)

// type UserNotFoundError struct {
// 	base.BaseError
// }

// func ErrUserNotFound() error {
// 	return &UserNotFoundError{
// 		BaseError: *base.NewBaseError(
// 			"User not found",
// 			404,
// 			"The specified user could not be found in the database",
// 		),
// 	}
// }

type DatabaseError struct {
	base.BaseError
}

func ErrDatabaseError(detail string) error {
	return &DatabaseError{
		BaseError: *base.NewBaseError(
			"Database error occurred",
			500,
			detail,
		),
	}
}

type ReconstructionError struct {
	base.BaseError
}

func ErrReconstructionError(detail string) error {
	return &ReconstructionError{
		BaseError: *base.NewBaseError(
			"Error reconstructing domain entity",
			500,
			detail,
		),
	}
}
