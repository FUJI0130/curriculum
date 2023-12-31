package sharedvo

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type CreatedAt time.Time

func NewCreatedAtByVal(createdAt time.Time) (CreatedAt, error) {
	if createdAt.IsZero() {
		return CreatedAt(time.Time{}), customerrors.NewUnprocessableEntityError("CreatedAt cannot be zero value")
	}
	if createdAt.After(time.Now()) {
		return CreatedAt(time.Time{}), customerrors.NewUnprocessableEntityError("CreatedAt is future")
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
