package sharedvo

import (
	"errors"
	"time"
)

type NewCreatedAtByVal time.Time
type CreatedAt time.Time

func NewCreatedAt(createdAt time.Time) (NewCreatedAtByVal, error) {
	if createdAt.IsZero() {
		return NewCreatedAtByVal(time.Time{}), errors.New("CreatedAt cannot be zero value")
	}
	if createdAt.After(time.Now()) {
		return NewCreatedAtByVal(time.Time{}), errors.New("CreatedAt cannot be future date")
	}
	NewCreatedAtValueObject := NewCreatedAtByVal(createdAt)
	return NewCreatedAtValueObject, nil
}

func (createdAt CreatedAt) DateTime() time.Time {
	CreatedAtDateTime := time.Time(createdAt)
	return CreatedAtDateTime
}

func (createdAt1 CreatedAt) Equal(createdAt2 CreatedAt) bool {

	return createdAt1 == createdAt2
}
