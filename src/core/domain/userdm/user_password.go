package userdm

import (
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

const (
	PasswordMinlength = 13
	PasswordMaxlength = 256
)

type UserPassword string

func NewUserPassword(userPassword string) (UserPassword, error) {
	count := utf8.RuneCountInString(userPassword)
	if userPassword == "" {
		return "", customerrors.NewUnprocessableEntityError("[NewUserPassword] password is empty")
	} else if PasswordMaxlength < count {
		return "", customerrors.NewUnprocessableEntityError("[NewUserPassword] password is too long")
	} else if count < PasswordMinlength {
		return "", customerrors.NewUnprocessableEntityError("[NewUserPassword] password is too short")
	}

	return UserPassword(userPassword), nil
}

func (password UserPassword) String() string {
	return string(password)
}

func (userPassword1 UserPassword) Equal(userPassword2 UserPassword) bool {
	return userPassword1 == userPassword2
}
