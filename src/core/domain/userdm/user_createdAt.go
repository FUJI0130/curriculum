package userdm

import (
	"errors"
	"time"
)

type UserCreatedAt struct {
	value time.Time
}

func NewUserCreatedAt(userCreatedAt time.Time) (*UserCreatedAt, error) {
	if userCreatedAt.IsZero() {
		return nil, errors.New("userCreatedAt cannot be zero value")
	}
	if userCreatedAt.After(time.Now()) {
		return nil, errors.New("userCreatedAt cannot be future date")
	}
	// userCreatedAt_VO := UserCreatedAt{value: userCreatedAt}
	return &UserCreatedAt{value: userCreatedAt}, nil
}

func (createdAt *UserCreatedAt) DateTime() time.Time {
	return createdAt.value
}

func (userCreatedAt1 *UserCreatedAt) Equal(userCreatedAt2 *UserCreatedAt) bool {
	return *userCreatedAt1 == *userCreatedAt2
}
