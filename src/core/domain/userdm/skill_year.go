package userdm

import "github.com/FUJI0130/curriculum/src/core/support/customerrors"

type SkillYear uint8

func NewSkillYear(yearsOfExperience uint8) (SkillYear, error) {
	// 経験年数が0以下であればエラーを返す
	if yearsOfExperience <= 0 {
		return 0, customerrors.NewUnprocessableEntityError("NewSkillYear:  SkillYear cannot be zero or negative value")
	}

	// 100年以上の経験は非現実的なので、このような上限も設定することができます。
	if yearsOfExperience > 100 {
		return 0, customerrors.NewUnprocessableEntityError("NewSkillYear: SkillYear should be less than or equal to 100")
	}

	skillsYearValueObject := SkillYear(yearsOfExperience)
	return skillsYearValueObject, nil
}

func (skillsYear SkillYear) Value() uint8 {
	return uint8(skillsYear)
}

func (skillsYear1 *SkillYear) Equal(skillsYear2 *SkillYear) bool {
	// 両方のポインタがnilの場合はtrueを返す
	if skillsYear1 == nil && skillsYear2 == nil {
		return true
	}

	// 片方のポインタだけがnilの場合はfalseを返す
	if skillsYear1 == nil || skillsYear2 == nil {
		return false
	}
	return *skillsYear1 == *skillsYear2
}
