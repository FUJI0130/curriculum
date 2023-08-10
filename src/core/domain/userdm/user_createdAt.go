package userdm

import (
	"errors"
	"time"
)

type UserCreatedAt time.Time

func NewUserCreatedAt(userCreatedAt time.Time) (*UserCreatedAt, error) {
	if userCreatedAt.IsZero() {
		return nil, errors.New("userCreatedAt cannot be zero value")
	}
	if userCreatedAt.After(time.Now()) {
		return nil, errors.New("userCreatedAt cannot be future date")
	}
	userCreatedAt_VO := UserCreatedAt(userCreatedAt)
	return &userCreatedAt_VO, nil
}

func (createdAt *UserCreatedAt) DateTime() time.Time {
	userCreatedAt_DateTime := time.Time(*createdAt)
	return userCreatedAt_DateTime
}

func (userCreatedAt1 *UserCreatedAt) Equal(userCreatedAt2 *UserCreatedAt) bool {
	return *userCreatedAt1 == *userCreatedAt2
}
