package userdm

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type UserPassword string

func NewUserPassword(userPassword string) (UserPassword, error) {
	count := utf8.RuneCountInString(userPassword)
	fmt.Println("Password Length:", count)
	if userPassword == "" {
		return "", errors.New("userPassword cannot be empty")
	} else if nameMaxlength < count {
		return "", errors.New("userPassword length over nameMaxlength")
	} else if count < 13 {
		return "", errors.New("userPassword length under 12")
	}

	userPasswordValueObject := UserPassword(userPassword)
	return userPasswordValueObject, nil
}

func (password UserPassword) String() string {
	return string(password)
}

func (userPassword1 UserPassword) Equal(userPassword2 UserPassword) bool {
	return userPassword1 == userPassword2
}
