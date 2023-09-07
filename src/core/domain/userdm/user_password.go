package userdm

import (
	"errors"
	"fmt"
)

type UserPassword string

func NewUserPassword(userPassword string) (UserPassword, error) {
	count := len([]rune(userPassword))
	fmt.Println("Password Length:", count)
	if userPassword == "" {
		return "", errors.New("userPassword cannot be empty")
	} else if 255 < count {
		return "", errors.New("userPassword length over 256")
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
