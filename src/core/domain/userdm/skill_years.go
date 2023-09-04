package userdm

import (
	"errors"
)

type SkillYears uint8

func NewSkillYears(yearsOfExperience uint8) (*SkillYears, error) {
	// 経験年数が0以下であればエラーを返す
	if yearsOfExperience <= 0 {
		return nil, errors.New("SkillYears cannot be zero or negative value")
	}

	// 100年以上の経験は非現実的なので、このような上限も設定することができます。
	if yearsOfExperience > 100 {
		return nil, errors.New("SkillYears is too long")
	}

	skillsYearsVO := SkillYears(yearsOfExperience)
	return &skillsYearsVO, nil
}

func (skillsYears SkillYears) Value() uint8 {
	return uint8(skillsYears)
}

func (skillsYears1 *SkillYears) Equal(skillsYears2 *SkillYears) bool {
	// 両方のポインタがnilの場合はtrueを返す
	if skillsYears1 == nil && skillsYears2 == nil {
		return true
	}

	// 片方のポインタだけがnilの場合はfalseを返す
	if skillsYears1 == nil || skillsYears2 == nil {
		return false
	}
	return *skillsYears1 == *skillsYears2
}
