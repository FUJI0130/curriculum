package userdm

import "errors"

type UserName string

func NewUserName(userName string) (*UserName, error) {
	if userName == "" {
		return nil, errors.New("userName cannot be empty")
	}
	userName_VO := UserName(userName)
	return &userName_VO, nil
}

func (name *UserName) String() string {
	return string(*name)
}

func (userName1 *UserName) Equal(userName2 *UserName) bool {
	return *userName1 == *userName2
}
