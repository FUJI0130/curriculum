package userdm

import (
	"errors"
	"time"
)

type UserUpdatedAt struct {
	value time.Time
}

func NewUserUpdatedAt(userUpdatedAt time.Time) (*UserUpdatedAt, error) {
	if userUpdatedAt.IsZero() {
		return nil, errors.New("userUpdatedAt cannot be zero value")
	}
	if userUpdatedAt.After(time.Now()) {
		return nil, errors.New("userUpdatedAt cannot be future date")
	}
	// userUpdatedAt_VO := UserUpdatedAt{value: userUpdatedAt}
	return &UserUpdatedAt{value: userUpdatedAt}, nil
}

func (updatedAt *UserUpdatedAt) DateTime() time.Time {
	return updatedAt.value
}

func (userUpdatedAt1 *UserUpdatedAt) Equal(userUpdatedAt2 *UserUpdatedAt) bool {
	return *userUpdatedAt1 == *userUpdatedAt2
}
