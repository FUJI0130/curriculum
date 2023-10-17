package userdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
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
		skill, err := NewSkill(param.TagID, user.ID(), param.Evaluation, param.Years, user.createdAt.DateTime(), user.updatedAt.DateTime())
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
