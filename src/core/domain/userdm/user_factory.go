package userdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

func GenWhenCreate(name string, email string, password string, profile string, skills []Skill, careers []Career) (*User, error) {
	user, err := NewUser(name, email, password, profile, skills, careers)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GenSkillWhenCreate(tagID string, userID string, evaluation uint8, years uint8) (*Skill, error) {
	skillID, err := NewSkillID()
	if err != nil {
		return nil, err
	}

	tID, err := tagdm.NewTagIDFromString(tagID)
	if err != nil {
		return nil, err
	}

	eval, err := NewSkillEvaluation(evaluation)
	if err != nil {
		return nil, err
	}

	y, err := NewSkillYear(years)
	if err != nil {
		return nil, err
	}

	skillCreatedAt := sharedvo.NewCreatedAt()

	skillUpdatedAt := sharedvo.NewUpdatedAt()

	return &Skill{
		id:         skillID,
		tagID:      tID,
		evaluation: eval,
		years:      y,
		createdAt:  skillCreatedAt,
		updatedAt:  skillUpdatedAt,
	}, nil
}

func GenCareerWhenCreate(detail string, adFrom time.Time, adTo time.Time, userID string) (*Career, error) {
	careerID, err := NewCareerID()
	if err != nil {
		return nil, err
	}

	if detail == "" {
		return nil, customerrors.NewUnprocessableEntityError("Career Detail is empty")
	}

	careerCreatedAt := sharedvo.NewCreatedAt()

	careerUpdatedAt := sharedvo.NewUpdatedAt()

	return &Career{
		id:        careerID,
		detail:    detail,
		adFrom:    adFrom,
		adTo:      adTo,
		createdAt: careerCreatedAt,
		updatedAt: careerUpdatedAt,
	}, nil
}

func GenSkillWhenUpdate(id SkillID, tagID tagdm.TagID, evaluation uint8, years uint8, createdAtVal time.Time) (*Skill, error) {

	// その他の値の設定
	eval, err := NewSkillEvaluation(evaluation)
	if err != nil {
		return nil, err
	}

	y, err := NewSkillYear(years)
	if err != nil {
		return nil, err
	}

	skillCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAtVal)
	if err != nil {
		return nil, err
	}

	skillUpdatedAt := sharedvo.NewUpdatedAt() // 現在の時刻を使用

	return &Skill{
		id:         id,
		tagID:      tagID,
		evaluation: eval,
		years:      y,
		createdAt:  skillCreatedAt,
		updatedAt:  skillUpdatedAt,
	}, nil
}

func GenCareerWhenUpdate(id CareerID, detail string, adFrom time.Time, adTo time.Time, createdAtVal time.Time) (*Career, error) {

	if detail == "" {
		return nil, customerrors.NewUnprocessableEntityError("Career Detail is empty")
	}

	careerCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAtVal)
	if err != nil {
		return nil, err
	}

	careerUpdatedAt := sharedvo.NewUpdatedAt() // 現在の時刻を使用

	return &Career{
		id:        id,
		detail:    detail,
		adFrom:    adFrom,
		adTo:      adTo,
		createdAt: careerCreatedAt,
		updatedAt: careerUpdatedAt,
	}, nil
}
