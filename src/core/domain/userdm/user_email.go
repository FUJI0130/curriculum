package userdm

import (
	"errors"
	"regexp"
)

type UserEmail string

func NewUserEmail(user_email string) (*UserEmail, error) {
	count := len([]rune(user_email))
	if user_email == "" {
		return nil, errors.New("userEmail cannot be empty")
	} else if 255 < count {
		return nil, errors.New("userEmail length over 256")
	}

	// メールアドレスの形式の正規表現
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(user_email) {
		return nil, errors.New("userEmail format is invalid")
	}

	userEmail_VO := UserEmail(user_email)
	return &userEmail_VO, nil
}

func (email UserEmail) String() string {
	userEmailString := string(email)
	return userEmailString
}

func (userEmail1 *UserEmail) Equal(userEmail2 *UserEmail) bool {
	return *userEmail1 == *userEmail2
}
