package userdm

import (
	"errors"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
)

type User struct {
	id        UserID
	name      string
	email     UserEmail
	password  UserPassword
	profile   string
	createdAt sharedvo.CreatedAt
	updatedAt sharedvo.UpdatedAt
}

func NewUser(name string, email string, password string, profile string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	if name == "" || password == "" {
		return nil, errors.New("name and password cannot be empty")
	}
	user_id, err := NewUserID()
	if err != nil {
		return nil, err
	}

	user_name := name
	//エラー処理ここに入れる
	count_name := len([]rune(user_name))
	if user_name == "" {
		return nil, errors.New("userName cannot be empty")
	} else if 255 < count_name {
		return nil, errors.New("userName length over 256")
	}

	user_email, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}

	user_password, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}

	user_profile := profile
	count_profile := len([]rune(profile))
	if user_profile == "" {
		return nil, errors.New("userProfile cannot be empty")
	} else if 255 < count_profile {
		return nil, errors.New("userProfile length over 256")
	}

	user_createdAt, err := sharedvo.NewCreatedAt(createdAt)
	if err != nil {
		return nil, err
	}

	user_updatedAt, err := sharedvo.NewUpdatedAt(updatedAt)
	if err != nil {
		return nil, err
	}

	return &User{
		id:        user_id,
		name:      user_name,
		email:     *user_email,
		password:  *user_password,
		profile:   user_profile,
		createdAt: *user_createdAt,
		updatedAt: *user_updatedAt,
	}, nil
}
