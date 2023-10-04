package userdm

import (
	"regexp"
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/customerrors"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm/constants"
)

type UserEmail string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewUserEmail(userEmail string) (UserEmail, error) {
	count := utf8.RuneCountInString(userEmail)
	if userEmail == "" {
		return "", customerrors.ErrUserEmailEmpty(nil, "NewUserEmail")
	} else if constants.EmailMaxlength < count {
		return "", customerrors.ErrUserEmailTooLong(nil, "NewUserEmail")
	}

	// メールアドレスの形式のチェック
	if !emailRegex.MatchString(userEmail) {
		return "", customerrors.ErrUserEmailInvalidFormat(nil, "NewUserEmail")
	}

	return UserEmail(userEmail), nil
}

func (email UserEmail) String() string {
	return string(email)
}

func (userEmail1 UserEmail) Equal(userEmail2 UserEmail) bool {
	return userEmail1 == userEmail2
}
