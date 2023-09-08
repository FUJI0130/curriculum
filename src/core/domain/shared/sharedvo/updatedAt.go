package sharedvo

import (
	"errors"
	"time"
)

var LastDuration time.Duration // 経過時間を保存するための変数

type UpdatedAt time.Time

func NewUpdatedAt(updatedAt time.Time) (UpdatedAt, error) {
	if updatedAt.IsZero() {
		return UpdatedAt(time.Time{}), errors.New("UpdatedAt cannot be zero value")
	}

	adjustTime := time.Now().Add(-1000 * time.Millisecond)

	if updatedAt.Before(adjustTime) {

		LastDuration = time.Since(adjustTime)
		return UpdatedAt(time.Time{}), errors.New("UpdatedAt cannot be past date")
	}

	return UpdatedAt(updatedAt), nil
}

func (updatedAt UpdatedAt) DateTime() time.Time {
	updatedAtDateTime := time.Time(updatedAt)
	return updatedAtDateTime
}

func (updatedAt *UpdatedAt) SetDateTime(newTime time.Time) error {
	// 制約を確認: 時間がゼロであってはならない
	if newTime.IsZero() {
		return errors.New("UpdatedAt cannot be zero value")
	}
	// 制約を確認: 過去の日付であってはならない
	if newTime.Before(time.Now()) {
		return errors.New("UpdatedAt cannot be past date")
	}
	// 制約を満たしていれば、新しい時間を設定
	*updatedAt = UpdatedAt(newTime)
	return nil
}

func (updatedAt1 UpdatedAt) Equal(updatedAt2 UpdatedAt) bool {

	return updatedAt1 == updatedAt2
}
