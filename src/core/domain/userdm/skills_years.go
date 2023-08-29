package userdm

import (
	"errors"
)

type SkillsYears int

func NewSkillsYears(yearsOfExperience int) (*SkillsYears, error) {
	// 経験年数が0以下であればエラーを返す
	if yearsOfExperience <= 0 {
		return nil, errors.New("SkillsYears cannot be zero or negative value")
	}

	// 100年以上の経験は非現実的なので、このような上限も設定することができます。
	if yearsOfExperience > 100 {
		return nil, errors.New("SkillsYears is too long")
	}

	skillsYearsVO := SkillsYears(yearsOfExperience)
	return &skillsYearsVO, nil
}

func (skillsYears SkillsYears) GetYearsOfExperience() int {
	return int(skillsYears)
}

func (skillsYears1 *SkillsYears) Equal(skillsYears2 *SkillsYears) bool {
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
