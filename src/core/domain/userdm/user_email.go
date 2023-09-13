package userdm

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

type UserEmail string

func emailRegex() string {
	return `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
}

func NewUserEmail(userEmail string) (UserEmail, error) {
	count := utf8.RuneCountInString(userEmail)
	if userEmail == "" {
		return "", errors.New("userEmail cannot be empty")
	} else if nameMaxlength < count {
		return "", errors.New("userEmail length over nameMaxlength")
	}

	// メールアドレスの形式の正規表現
	// var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex())
	if !re.MatchString(userEmail) {
		return "", errors.New("userEmail format is invalid")
	}

	return UserEmail(userEmail), nil
}

func (email UserEmail) String() string {
	return string(email)
}

func (userEmail1 UserEmail) Equal(userEmail2 UserEmail) bool {
	return userEmail1 == userEmail2
}
