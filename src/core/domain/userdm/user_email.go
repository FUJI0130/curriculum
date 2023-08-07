package userdm

import "errors"

type UserEmail string

func NewUserEmail(userEmail string) (*UserEmail, error) {
	if userEmail == "" {
		return nil, errors.New("userEmail cannot be empty")
	}
	userEmail_VO := UserEmail(userEmail)
	return &userEmail_VO, nil
}

func (email *UserEmail) String() string {
	return string(*email)
}

func (userEmail1 *UserEmail) Equal(userEmail2 *UserEmail) bool {
	return *userEmail1 == *userEmail2
}
