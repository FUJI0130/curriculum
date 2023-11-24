package userdm

import (
	"time"

	"github.com/cockroachdb/errors"

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

func GenWhenUpdate(updatedUser *User, updatedSkills []*Skill, updatedCareers []*Career) (*User, error) {
	if updatedUser == nil {
		return nil, errors.New("updated user is nil")
	}

	newUser := *updatedUser

	newSkills := make([]Skill, len(updatedSkills))
	for i, s := range updatedSkills {
		if s != nil {
			newSkills[i] = *s
		}
	}
	newUser.skills = newSkills

	newCareers := make([]Career, len(updatedCareers))
	for i, c := range updatedCareers {
		if c != nil {
			newCareers[i] = *c
		}
	}
	newUser.careers = newCareers

	newUser.updatedAt = sharedvo.NewUpdatedAt()

	return &newUser, nil
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
	skillId, err := NewSkillIDFromString(id)
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

func GenCareerWhenUpdate(id string, detail string, adFrom time.Time, adTo time.Time, userID string, createdAt time.Time, updatedAt time.Time) (*Career, error) {
	careerId, err := NewCareerIDFromString(id)
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
