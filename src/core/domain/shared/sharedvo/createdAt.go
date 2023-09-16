package sharedvo

import (
	"errors"
	"time"
)

type CreatedAt time.Time

func NewCreatedAtByVal(createdAt time.Time) (CreatedAt, error) {
	if createdAt.IsZero() {
		return CreatedAt(time.Time{}), errors.New("CreatedAt cannot be zero value")
	}
	if createdAt.After(time.Now()) {
		return CreatedAt(time.Time{}), errors.New("CreatedAt cannot be future date")
	}
	return CreatedAt(createdAt), nil
}

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func (createdAt CreatedAt) DateTime() time.Time {
	return time.Time(createdAt)
}

func (createdAt1 CreatedAt) Equal(createdAt2 CreatedAt) bool {
	return createdAt1 == createdAt2
}
