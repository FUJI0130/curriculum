package sharedvo

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type UpdatedAt time.Time

func NewUpdatedAtByVal(updatedAt time.Time) (UpdatedAt, error) {
	if updatedAt.IsZero() {
		return UpdatedAt(time.Time{}), customerrors.NewUnprocessableEntityError("UpdatedAt is empty")
	}
	return UpdatedAt(updatedAt), nil
}

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now())
}

func (updatedAt UpdatedAt) DateTime() time.Time {
	return time.Time(updatedAt)
}

func (updatedAt1 UpdatedAt) Equal(updatedAt2 UpdatedAt) bool {
	return updatedAt1 == updatedAt2
}
