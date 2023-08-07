package userdm

import "errors"

type UserPassword string

func NewUserPassword(userPassword string) (*UserPassword, error) {
	if userPassword == "" {
		return nil, errors.New("userPassword cannot be empty")
	}
	userPassword_VO := UserPassword(userPassword)
	return &userPassword_VO, nil
}

func (password *UserPassword) String() string {
	return string(*password)
}

func (userPassword1 *UserPassword) Equal(userPassword2 *UserPassword) bool {
	return *userPassword1 == *userPassword2
}
