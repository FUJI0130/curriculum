package userdm

import "errors"

type UserName string

func NewUserName(userName string) (*UserName, error) {
	count := len([]rune(userName))
	if userName == "" {
		return nil, errors.New("userName cannot be empty")
	} else if 255 < count {
		return nil, errors.New("userName length over 256")
	}
	userName_VO := UserName(userName)
	return &userName_VO, nil
}

func (name *UserName) String() string {
	userNameString := string(*name)
	return userNameString
}

func (userName1 *UserName) Equal(userName2 *UserName) bool {
	return *userName1 == *userName2
}
