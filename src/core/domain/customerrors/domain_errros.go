package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
)

type ReconstructionError struct {
	base.BaseError
}

func ErrReconstructionError(cause error, customMsg ...string) *base.BaseError {
	fullMessage := "Error reconstructing domain entity"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return base.NewBaseError(
		fullMessage,
		errorcodes.InternalServerError,
		cause,
	)
}
