package userdm

import "errors"

type UserEmail struct {
	userEmail string
}

func NewUserEmail(user_email string) (*UserEmail, error) {
	if user_email == "" {
		return nil, errors.New("userEmail cannot be empty")
	}
	userEmail_VO := UserEmail{userEmail: user_email}
	return &userEmail_VO, nil
}

func (email *UserEmail) String() string {
	return string(*&email.userEmail)
}

func (userEmail1 *UserEmail) Equal(userEmail2 *UserEmail) bool {
	return *userEmail1 == *userEmail2
}
