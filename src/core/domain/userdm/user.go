package userdm

import (
	"errors"
	"fmt"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
)

type User struct {
	id        UserID             `db:"id"`
	name      string             `db:"name"`
	email     UserEmail          `db:"email"`
	password  UserPassword       `db:"password"`
	profile   string             `db:"profile"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

var ErrUserNotFound = errors.New("user not found")

func NewUser(name string, email string, password string, profile string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	if name == "" || password == "" {
		return nil, errors.New("name and password cannot be empty")
	}
	userID, err := NewUserID()
	if err != nil {
		return nil, err
	}

	userName := name
	//エラー処理ここに入れる
	countName := len([]rune(userName))
	if userName == "" {
		return nil, errors.New("userName cannot be empty")
	} else if 255 < countName {
		return nil, errors.New("userName length over 256")
	}

	userEmail, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}

	userPassword, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}

	userProfile := profile
	count_profile := len([]rune(profile))
	if userProfile == "" {
		return nil, errors.New("userProfile cannot be empty")
	} else if 255 < count_profile {
		return nil, errors.New("userProfile length over 256")
	}

	userCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}

	userUpdatedAt, err := sharedvo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		fmt.Printf("NewUser Time taken for updatedAt.Before(time.Now()): %v\n", sharedvo.LastDuration)
		return nil, err
	}

	return &User{
		id:        userID,
		name:      userName,
		email:     userEmail,
		password:  userPassword,
		profile:   userProfile,
		createdAt: userCreatedAt,
		updatedAt: userUpdatedAt,
	}, nil
}

func ReconstructUser(id UserID, name string, email string, password string, profile string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	// 名前、Eメール、パスワードなどのエラーチェックをここで行うことができます
	userName := name
	//エラー処理ここに入れる
	countName := len([]rune(userName))
	if userName == "" {
		return nil, errors.New("userName cannot be empty")
	} else if 255 < countName {
		return nil, errors.New("userName length over 256")
	}

	userEmail, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}

	userPassword, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}

	userProfile := profile
	count_profile := len([]rune(profile))
	if userProfile == "" {
		return nil, errors.New("userProfile cannot be empty")
	} else if 255 < count_profile {
		return nil, errors.New("userProfile length over 256")
	}

	userCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}

	userUpdatedAt, err := sharedvo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		fmt.Printf("ReconstructUser Time taken for updatedAt.Before(time.Now()): %v\n", sharedvo.LastDuration)
		return nil, err
	}
	return &User{
		id:        id,
		name:      userName,
		email:     userEmail,
		password:  userPassword,
		profile:   userProfile,
		createdAt: userCreatedAt,
		updatedAt: userUpdatedAt,
	}, nil
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() UserEmail {
	return u.email
}

func (u *User) Password() UserPassword {
	return u.password
}

func (u *User) Profile() string {
	return u.profile
}

func (u *User) CreatedAt() sharedvo.CreatedAt {
	return u.createdAt
}

func (u *User) UpdatedAt() sharedvo.UpdatedAt {
	return u.updatedAt
}
