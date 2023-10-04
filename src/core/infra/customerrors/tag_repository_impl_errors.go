package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/common/base"
	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
)

type TagNotFoundError struct {
	base.BaseError
}

func ErrTagNotFound(cause error, customMsg ...string) error {
	defaultMessage := "Tag not found"
	detail := "The specified tag could not be found in the database"
	if len(customMsg) > 0 && customMsg[0] != "" {
		detail = customMsg[0]
	} else if cause != nil {
		detail = fmt.Sprintf("%s: %v", detail, cause)
	}
	return &TagNotFoundError{
		BaseError: *base.NewBaseError(defaultMessage, errorcodes.NotFound, cause),
	}
}
