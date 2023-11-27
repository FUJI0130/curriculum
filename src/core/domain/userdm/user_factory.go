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

func GenSkillWhenCreate(tagID string, userID string, evaluation uint8, years uint8, createdAt time.Time, updatedAt time.Time) (*Skill, error) {
	skillId, err := NewSkillID()
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

	skillCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}

	skillUpdatedAt, err := sharedvo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		return nil, err
	}

	return &Skill{
		id:         skillId,
		tagID:      tID,
		evaluation: eval,
		years:      y,
		createdAt:  skillCreatedAt,
		updatedAt:  skillUpdatedAt,
	}, nil
}

func GenCareerWhenCreate(detail string, adFrom time.Time, adTo time.Time, userID string, createdAt time.Time, updatedAt time.Time) (*Career, error) {
	careerId, err := NewCareerID()
	if err != nil {
		return nil, err
	}

	if detail == "" {
		return nil, customerrors.NewUnprocessableEntityError("Career Detail is empty")
	}

	careerCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}

	careerUpdatedAt, err := sharedvo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		return nil, err
	}

	return &Career{
		id:        careerId,
		detail:    detail,
		adFrom:    adFrom,
		adTo:      adTo,
		createdAt: careerCreatedAt,
		updatedAt: careerUpdatedAt,
	}, nil
}

func GenSkillWhenUpdate(id string, tagIdStr string, evaluation uint8, years uint8, createdAtVal time.Time) (*Skill, error) {
	var skillId SkillID
	var err error

	// SkillIDの生成またはパース
	if id != "" {
		skillId, err = NewSkillIDFromString(id)
		if err != nil {
			return nil, err
		}
	} else {
		skillId, err = NewSkillID()
		if err != nil {
			return nil, err
		}
	}

	// TagIDの生成
	tagId, err := tagdm.NewTagIDFromString(tagIdStr)
	if err != nil {
		return nil, err
	}

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
		id:         skillId,
		tagID:      tagId,
		evaluation: eval,
		years:      y,
		createdAt:  skillCreatedAt,
		updatedAt:  skillUpdatedAt,
	}, nil
}

func GenCareerWhenUpdate(id string, detail string, adFrom time.Time, adTo time.Time, createdAtVal time.Time) (*Career, error) {
	var careerId CareerID
	var err error

	if id != "" {
		careerId, err = NewCareerIDFromString(id)
		if err != nil {
			return nil, err
		}
	} else {
		careerId, err = NewCareerID()
		if err != nil {
			return nil, err
		}
	}

	if detail == "" {
		return nil, customerrors.NewUnprocessableEntityError("Career Detail is empty")
	}

	careerCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAtVal)
	if err != nil {
		return nil, err
	}

	careerUpdatedAt := sharedvo.NewUpdatedAt() // 現在の時刻を使用

	return &Career{
		id:        careerId,
		detail:    detail,
		adFrom:    adFrom,
		adTo:      adTo,
		createdAt: careerCreatedAt,
		updatedAt: careerUpdatedAt,
	}, nil
}
