package sharedvo

import (
	"errors"
	"time"
)

type CreatedAt time.Time

func NewCreatedAt(createdAt time.Time) (CreatedAt, error) {
	if createdAt.IsZero() {
		return CreatedAt(time.Time{}), errors.New("CreatedAt cannot be zero value")
	}
	if createdAt.After(time.Now()) {
		return CreatedAt(time.Time{}), errors.New("CreatedAt cannot be future date")
	}
	CreatedAtValueObject := CreatedAt(createdAt)
	return CreatedAtValueObject, nil
}

func (createdAt CreatedAt) DateTime() time.Time {
	CreatedAtDateTime := time.Time(createdAt)
	return CreatedAtDateTime
}

func (createdAt1 CreatedAt) Equal(createdAt2 CreatedAt) bool {

	return createdAt1 == createdAt2
}
