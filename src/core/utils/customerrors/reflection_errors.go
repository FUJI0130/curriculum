package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/FUJI0130/curriculum/src/core/support/errorcodes"
)

type ReflectionError struct {
	base.BaseError
}

func ErrStructToMap(cause error, customMsg ...string) error {
	fullMessage := "Failed to convert struct to map"
	if len(customMsg) > 0 && customMsg[0] != "" {
		fullMessage = fmt.Sprintf("%s: %s", fullMessage, customMsg[0])
	} else if cause != nil {
		fullMessage = fmt.Sprintf("%s: %v", fullMessage, cause)
	}
	return &ReflectionError{
		BaseError: *base.NewBaseError(fullMessage, errorcodes.InternalServerError, cause),
	}
}
