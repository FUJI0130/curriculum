package userdm

import (
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm/constants"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type UserPassword string

func NewUserPassword(userPassword string) (UserPassword, error) {
	count := utf8.RuneCountInString(userPassword)
	if userPassword == "" {
		return "", customerrors.NewUnprocessableEntityError("NewUserPassword password is empty")
	} else if constants.PasswordMaxlength < count {
		return "", customerrors.NewUnprocessableEntityError("NewUserPassword password is too long")
	} else if count < constants.PasswordMinlength {
		return "", customerrors.NewUnprocessableEntityError("NewUserPassword password is too short")
	}

	return UserPassword(userPassword), nil
}

func (password UserPassword) String() string {
	return string(password)
}

func (userPassword1 UserPassword) Equal(userPassword2 UserPassword) bool {
	return userPassword1 == userPassword2
}
