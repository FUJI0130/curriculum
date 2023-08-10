package sharedvo

import (
	"errors"
	"time"
)

type CreatedAt time.Time

func NewCreatedAt(createdAt time.Time) (*CreatedAt, error) {
	if createdAt.IsZero() {
		return nil, errors.New("CreatedAt cannot be zero value")
	}
	if createdAt.After(time.Now()) {
		return nil, errors.New("CreatedAt cannot be future date")
	}
	CreatedAt_VO := CreatedAt(createdAt)
	return &CreatedAt_VO, nil
}

func (createdAt *CreatedAt) DateTime() time.Time {
	CreatedAt_DateTime := time.Time(*createdAt)
	return CreatedAt_DateTime
}

func (createdAt1 *CreatedAt) Equal(createdAt2 *CreatedAt) bool {
	return *createdAt1 == *createdAt2
}
