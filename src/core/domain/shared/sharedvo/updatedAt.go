package sharedvo

import (
	"errors"
	"time"
)

var LastDuration time.Duration // 経過時間を保存するための変数

type UpdatedAt time.Time

// NewUpdatedAtByVal は指定した時刻での UpdatedAt を生成する
func NewUpdatedAtByVal(updatedAt time.Time) (UpdatedAt, error) {
	if updatedAt.IsZero() {
		return UpdatedAt(time.Time{}), errors.New("UpdatedAt cannot be zero value")
	}

	// 制約を確認: 過去の日付であってはならない
	if updatedAt.Before(time.Now().Add(-time.Second)) {
		return UpdatedAt(time.Time{}), errors.New("UpdatedAt cannot be past date")
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
