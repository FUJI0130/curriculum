package sharedvo

import (
	"errors"
	"time"
)

type UpdatedAt time.Time

func NewUpdatedAt(updatedAt time.Time) (*UpdatedAt, error) {
	if updatedAt.IsZero() {
		return nil, errors.New("UpdatedAt cannot be zero value")
	}
	if updatedAt.After(time.Now()) {
		return nil, errors.New("UpdatedAt cannot be future date")
	}
	UpdatedAt_VO := UpdatedAt(updatedAt)
	return &UpdatedAt_VO, nil
}

func (updatedAt *UpdatedAt) DateTime() time.Time {
	UpdatedAt_DateTime := time.Time(*updatedAt)
	return UpdatedAt_DateTime
}

func (updatedAt1 *UpdatedAt) Equal(updatedAt2 *UpdatedAt) bool {
	return *updatedAt1 == *updatedAt2
}
