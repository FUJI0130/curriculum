package userdm

import (
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/customerrors"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm/constants"
)

type UserPassword string

func NewUserPassword(userPassword string) (UserPassword, error) {
	count := utf8.RuneCountInString(userPassword)
	if userPassword == "" {
		return "", customerrors.ErrUserPasswordEmpty()
	} else if constants.PasswordMaxlength < count {
		return "", customerrors.ErrUserPasswordTooLong()
	} else if count < constants.PasswordMinlength {
		return "", customerrors.ErrUserPasswordTooShort()
	}

	return UserPassword(userPassword), nil
}

func (password UserPassword) String() string {
	return string(password)
}

func (userPassword1 UserPassword) Equal(userPassword2 UserPassword) bool {
	return userPassword1 == userPassword2
}
