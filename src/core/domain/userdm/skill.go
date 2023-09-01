package userdm

import (
	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
)

type Skill struct {
	id         SkillID            `db:"id"`
	tagId      tagdm.TagID        `db:"tag_id"`
	userId     UserID             `db:"user_id"`
	evaluation SkillEvaluation    `db:"evaluation"`
	years      SkillYears         `db:"years"`
	createdAt  sharedvo.CreatedAt `db:"created_at"`
	updatedAt  sharedvo.UpdatedAt `db:"updated_at"`
}

func NewSkill(tagId tagdm.TagID, userId UserID, evaluation uint8, years uint8) (*Skill, error) {
	eval, err := NewSkillEvaluation(evaluation)
	if err != nil {
		return nil, err
	}

	y, err := NewSkillYears(years)
	if err != nil {
		return nil, err
	}

	return &Skill{
		tagId:      tagId,
		userId:     userId,
		evaluation: *eval,
		years:      *y,
	}, nil
}
func (s *Skill) ID() SkillID {
	return s.id
}
func (s *Skill) TagID() tagdm.TagID {
	return s.tagId
}
func (s *Skill) UserID() UserID {
	return s.userId
}

func (s *Skill) Evaluation() SkillEvaluation {
	return s.evaluation
}

func (s *Skill) Years() SkillYears {
	return s.years
}
func (s *Skill) CreatedAt() sharedvo.CreatedAt {
	return s.createdAt
}

func (s *Skill) UpdatedAt() sharedvo.UpdatedAt {
	return s.updatedAt
}
