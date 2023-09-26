package sharedvo

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/customerrors"
)

type UpdatedAt time.Time

// NewUpdatedAtByVal は指定した時刻での UpdatedAt を生成する
func NewUpdatedAtByVal(updatedAt time.Time) (UpdatedAt, error) {
	if updatedAt.IsZero() {
		return UpdatedAt(time.Time{}), customerrors.ErrUpdatedAtZeroValue()
	}

	return UpdatedAt(updatedAt), nil
}

// NewUpdatedAt は現在の時刻での UpdatedAt を生成する
func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now())
}

func (updatedAt UpdatedAt) DateTime() time.Time {
	return time.Time(updatedAt)
}

func (updatedAt1 UpdatedAt) Equal(updatedAt2 UpdatedAt) bool {
	return updatedAt1 == updatedAt2
}
