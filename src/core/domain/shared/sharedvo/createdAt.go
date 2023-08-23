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

func (createdAt CreatedAt) DateTime() time.Time {
	CreatedAtDateTime := time.Time(createdAt)
	return CreatedAtDateTime
}

func (createdAt1 *CreatedAt) Equal(createdAt2 *CreatedAt) bool {
	// 両方のポインタがnilの場合はtrueを返す
	if createdAt1 == nil && createdAt2 == nil {
		return true
	}

	// 片方のポインタだけがnilの場合はfalseを返す
	if createdAt1 == nil || createdAt2 == nil {
		return false
	}
	return *createdAt1 == *createdAt2
}
