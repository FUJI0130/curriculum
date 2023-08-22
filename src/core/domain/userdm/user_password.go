package userdm

import (
	"errors"
	"fmt"
)

type UserPassword string

func NewUserPassword(userPassword string) (*UserPassword, error) {
	count := len([]rune(userPassword))
	fmt.Println("Password Length:", count)
	if userPassword == "" {
		return nil, errors.New("userPassword cannot be empty")
	} else if 255 < count {
		return nil, errors.New("userPassword length over 256")
	} else if count < 13 {
		return nil, errors.New("userPassword length under 12")
	}

	userPassword_VO := UserPassword(userPassword)
	return &userPassword_VO, nil
}

func (password UserPassword) String() string {
	userPasswordString := string(password)
	return userPasswordString
}

func (userPassword1 *UserPassword) Equal(userPassword2 *UserPassword) bool {
	return *userPassword1 == *userPassword2
}
