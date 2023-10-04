package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
)

type ValidatorError struct {
	base.BaseError
}

func ErrValidateKeysAgainstStruct(cause error, customMsg ...string) error {
	fullMessage := "Validation failed"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &ValidatorError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.BadRequest, cause),
	}
}
