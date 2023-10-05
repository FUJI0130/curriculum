package userdm

import (
	"regexp"
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm/constants"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type UserEmail string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewUserEmail(userEmail string) (UserEmail, error) {
	count := utf8.RuneCountInString(userEmail)
	if userEmail == "" {
		return "", customerrors.NewUnprocessableEntityError("NewUserEmail email is empty")
	} else if constants.EmailMaxlength < count {
		return "", customerrors.NewUnprocessableEntityError("NewUserEmail email is too long")
	}

	// メールアドレスの形式のチェック
	if !emailRegex.MatchString(userEmail) {
		return "", customerrors.NewUnprocessableEntityError("NewUserEmail email is invalid")
	}

	return UserEmail(userEmail), nil
}

func (email UserEmail) String() string {
	return string(email)
}

func (userEmail1 UserEmail) Equal(userEmail2 UserEmail) bool {
	return userEmail1 == userEmail2
}
