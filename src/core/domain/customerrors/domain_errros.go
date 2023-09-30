package customerrors

import (
	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
)

type ReconstructionError struct {
	base.BaseError
}

func ErrReconstructionError(detail string) error {
	return &ReconstructionError{
		BaseError: *base.NewBaseError(
			"Error reconstructing domain entity",
			errorcodes.InternalServerError,
			detail,
		),
	}
}
