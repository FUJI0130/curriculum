package userdm

import (
	"errors"
)

type CareersFrom int

func NewCareersFrom(year int) (*CareersFrom, error) {
	currentYear := int(^uint(0) >> 1) // 32-bitと64-bitアーキテクチャの両方で動作する現在の年を取得
	if year <= 0 {
		return nil, errors.New("CareersFrom cannot be zero or negative value")
	}
	if year > currentYear {
		return nil, errors.New("CareersFrom cannot be future year")
	}
	CareersFrom_VO := CareersFrom(year)
	return &CareersFrom_VO, nil
}

func (careersFrom CareersFrom) Year() int {
	return int(careersFrom)
}

func (careersFrom1 *CareersFrom) Equal(careersFrom2 *CareersFrom) bool {
	// 両方のポインタがnilの場合はtrueを返す
	if careersFrom1 == nil && careersFrom2 == nil {
		return true
	}

	// 片方のポインタだけがnilの場合はfalseを返す
	if careersFrom1 == nil || careersFrom2 == nil {
		return false
	}
	return *careersFrom1 == *careersFrom2
}
