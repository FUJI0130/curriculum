package userdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type Skill struct {
	id         SkillID            `db:"id"`
	tagID      tagdm.TagID        `db:"tag_id"`
	userID     UserID             `db:"user_id"`
	evaluation SkillEvaluation    `db:"evaluation"`
	years      SkillYear          `db:"years"`
	createdAt  sharedvo.CreatedAt `db:"created_at"`
	updatedAt  sharedvo.UpdatedAt `db:"updated_at"`
}

func NewSkill(tagID tagdm.TagID, userID UserID, evaluation uint8, years uint8, createdAt time.Time, updatedAt time.Time) (*Skill, error) {
	eval, err := NewSkillEvaluation(evaluation)
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityErrorf(err, "evaluation value: %d", evaluation)
	}

	y, err := NewSkillYear(years)
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityErrorf(err, "years value: %d", years)
	}

	skillCreatedAt := sharedvo.NewCreatedAt()
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityError(err, "createdAt is invalid")
	}

	skillUpdatedAt := sharedvo.NewUpdatedAt()
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityError(err, "updatedAt is invalid")
	}
	return &Skill{
		tagID:      tagID,
		userID:     userID,
		evaluation: eval,
		years:      y,
		createdAt:  skillCreatedAt,
		updatedAt:  skillUpdatedAt,
	}, nil
}
func (s *Skill) ID() SkillID {
	return s.id
}
func (s *Skill) TagID() tagdm.TagID {
	return s.tagID
}
func (s *Skill) UserID() UserID {
	return s.userID
}

func (s *Skill) Evaluation() SkillEvaluation {
	return s.evaluation
}

func (s *Skill) Year() SkillYear {
	return s.years
}
func (s *Skill) CreatedAt() sharedvo.CreatedAt {
	return s.createdAt
}

func (s *Skill) UpdatedAt() sharedvo.UpdatedAt {
	return s.updatedAt
}
