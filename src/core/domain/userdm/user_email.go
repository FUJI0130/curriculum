package userdm

import (
	"errors"
	"regexp"
)

type UserEmail string

func NewUserEmail(userEmail string) (*UserEmail, error) {
	count := len([]rune(userEmail))
	if userEmail == "" {
		return nil, errors.New("userEmail cannot be empty")
	} else if 255 < count {
		return nil, errors.New("userEmail length over 256")
	}

	// メールアドレスの形式の正規表現
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(userEmail) {
		return nil, errors.New("userEmail format is invalid")
	}

	userEmailVO := UserEmail(userEmail)
	return &userEmailVO, nil
}

func (email UserEmail) String() string {
	userEmailString := string(email)
	return userEmailString
}

func (userEmail1 *UserEmail) Equal(userEmail2 *UserEmail) bool {
	return *userEmail1 == *userEmail2
}
