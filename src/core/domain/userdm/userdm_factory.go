package userdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type SkillParam struct {
	TagID      tagdm.TagID
	TagName    string
	Evaluation uint8
	Years      uint8
}

type CareerParam struct {
	Detail string
	AdFrom time.Time
	AdTo   time.Time
}

func GenWhenCreate(name string, email string, password string, profile string, skillParams []SkillParam, careerParams []CareerParam) (*UserDomain, error) {
	user, err := NewUser(name, email, password, profile)
	if err != nil {
		return nil, err
	}

	skills := make([]*Skill, 0, len(skillParams))
	for _, param := range skillParams {
		skill, err := NewSkill(param.TagID, user.ID(), param.Evaluation, param.Years)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	careers := make([]*Career, 0, len(careerParams))
	for _, param := range careerParams {
		career, err := NewCareer(param.Detail, param.AdFrom, param.AdTo, user.ID())
		if err != nil {
			return nil, err
		}
		careers = append(careers, career)
	}

	return NewUserDomain(user, skills, careers), nil
}

func GenSkillWhenCreate(tagID string, userID string, evaluation uint8, years uint8, createdAt time.Time, updatedAt time.Time) (*Skill, error) {
	skillId, err := NewSkillID() // ID文字列からSkillIDを再構築する関数を想定
	if err != nil {
		return nil, err
	}

	tID, err := tagdm.NewTagIDFromString(tagID) // TagIDを再構築する関数を想定
	if err != nil {
		return nil, err
	}

	uID, err := NewUserIDFromString(userID) // UserIDを再構築する関数を想定
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
		userID:     uID,
		evaluation: eval,
		years:      y,
		createdAt:  skillCreatedAt,
		updatedAt:  skillUpdatedAt,
	}, nil
}

func GenCareerWhenCreate(detail string, adFrom time.Time, adTo time.Time, userID string, createdAt time.Time, updatedAt time.Time) (*Career, error) {
	careerId, err := NewCareerID() // Assuming NewCareerIDFromString function exists
	if err != nil {
		return nil, err
	}

	uID, err := NewUserIDFromString(userID) // UserID reconstruction function, as per Skill example
	if err != nil {
		return nil, err
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
		userID:    uID,
		createdAt: careerCreatedAt,
		updatedAt: careerUpdatedAt,
	}, nil
}

func GenUserWhenUpdate(id string, name string, email string, password string, profile string, createdAt time.Time) (*User, error) {
	userId, err := NewUserIDFromString(id)
	if err != nil {
		return nil, err
	}
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("name is empty")
	}
	userEmail, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}

	userPassword, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}
	userCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}

	userUpdatedAt, err := sharedvo.NewUpdatedAtByVal(time.Now())
	if err != nil {
		return nil, err
	}
	return &User{
		id:        userId,
		name:      name,
		email:     userEmail,
		password:  userPassword,
		profile:   profile,
		createdAt: userCreatedAt,
		updatedAt: userUpdatedAt,
	}, nil
}

func GenSkillWhenUpdate(id string, tagID string, userID string, evaluation uint8, years uint8, createdAt time.Time, updatedAt time.Time) (*Skill, error) {
	skillId, err := NewSkillIDFromString(id) // ID文字列からSkillIDを再構築する関数を想定
	if err != nil {
		return nil, err
	}

	tID, err := tagdm.NewTagIDFromString(tagID) // TagIDを再構築する関数を想定
	if err != nil {
		return nil, err
	}

	uID, err := NewUserIDFromString(userID) // UserIDを再構築する関数を想定
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
		userID:     uID,
		evaluation: eval,
		years:      y,
		createdAt:  skillCreatedAt,
		updatedAt:  skillUpdatedAt,
	}, nil
}

func GenCareerWhenUpdate(id string, detail string, adFrom time.Time, adTo time.Time, userID string, createdAt time.Time, updatedAt time.Time) (*Career, error) {
	careerId, err := NewCareerIDFromString(id) // Assuming NewCareerIDFromString function exists
	if err != nil {
		return nil, err
	}

	uID, err := NewUserIDFromString(userID) // UserID reconstruction function, as per Skill example
	if err != nil {
		return nil, err
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
		userID:    uID,
		createdAt: careerCreatedAt,
		updatedAt: careerUpdatedAt,
	}, nil
}
