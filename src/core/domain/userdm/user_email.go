package userdm

import "errors"

type UserEmail string

func NewUserEmail(user_email string) (*UserEmail, error) {
	if user_email == "" {
		return nil, errors.New("userEmail cannot be empty")
	}
	userEmail_VO := UserEmail(user_email)
	return &userEmail_VO, nil
}

func (email *UserEmail) String() string {
	return email.String()
}

func (userEmail1 *UserEmail) Equal(userEmail2 *UserEmail) bool {
	return *userEmail1 == *userEmail2
}
