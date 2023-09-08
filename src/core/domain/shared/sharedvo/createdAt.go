package sharedvo

import (
	"errors"
	"time"
)

type CreatedAt time.Time

// NewCreatedAtByVal は指定した時刻での CreatedAt を生成する
func NewCreatedAtByVal(createdAt time.Time) (CreatedAt, error) {
	if createdAt.IsZero() {
		return CreatedAt(time.Time{}), errors.New("CreatedAt cannot be zero value")
	}
	if createdAt.After(time.Now()) {
		return CreatedAt(time.Time{}), errors.New("CreatedAt cannot be future date")
	}
	return CreatedAt(createdAt), nil
}

// NewCreatedAt は現在の時刻での CreatedAt を生成する
func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func (createdAt CreatedAt) DateTime() time.Time {
	return time.Time(createdAt)
}

func (createdAt1 CreatedAt) Equal(createdAt2 CreatedAt) bool {
	return createdAt1 == createdAt2
}
