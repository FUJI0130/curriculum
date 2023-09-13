package userdm

import (
	"errors"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
)

const nameMaxlength = 256

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
