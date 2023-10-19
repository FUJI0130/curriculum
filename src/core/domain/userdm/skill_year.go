package userdm

import "github.com/FUJI0130/curriculum/src/core/support/customerrors"

type SkillYear uint8

func NewSkillYear(yearsOfExperience uint8) (SkillYear, error) {
	if yearsOfExperience <= 0 {
		return 0, customerrors.NewUnprocessableEntityError("SkillYear cannot be zero or negative value")
	}

	if yearsOfExperience > 100 {
		return 0, customerrors.NewUnprocessableEntityError("SkillYear should be less than or equal to 100")
	}

	skillsYearValueObject := SkillYear(yearsOfExperience)
	return skillsYearValueObject, nil
}

func (skillsYear SkillYear) Value() uint8 {
	return uint8(skillsYear)
}

func (skillsYear1 *SkillYear) Equal(skillsYear2 *SkillYear) bool {
	if skillsYear1 == nil && skillsYear2 == nil {
		return true
	}

	if skillsYear1 == nil || skillsYear2 == nil {
		return false
	}
	return *skillsYear1 == *skillsYear2
}
