package userdm

import (
	"errors"
)

type CareersTo int

func NewCareersTo(year int) (*CareersTo, error) {
	currentYear := int(^uint(0) >> 1) // 32-bitと64-bitアーキテクチャの両方で動作する現在の年を取得
	if year <= 0 {
		return nil, errors.New("CareersTo cannot be zero or negative value")
	}
	if year > currentYear {
		return nil, errors.New("CareersTo cannot be future year")
	}
	CareersTo_VO := CareersTo(year)
	return &CareersTo_VO, nil
}

func (careersTo CareersTo) Year() int {
	return int(careersTo)
}

func (careersTo1 *CareersTo) Equal(careersTo2 *CareersTo) bool {
	// 両方のポインタがnilの場合はtrueを返す
	if careersTo1 == nil && careersTo2 == nil {
		return true
	}

	// 片方のポインタだけがnilの場合はfalseを返す
	if careersTo1 == nil || careersTo2 == nil {
		return false
	}
	return *careersTo1 == *careersTo2
}
