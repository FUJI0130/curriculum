package userdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
)

type Skill struct {
	id         SkillID            `db:"id"`
	tagID      tagdm.TagID        `db:"tag_id"`
	evaluation SkillEvaluation    `db:"evaluation"`
	years      SkillYear          `db:"years"`
	createdAt  sharedvo.CreatedAt `db:"created_at"`
	updatedAt  sharedvo.UpdatedAt `db:"updated_at"`
}

func NewSkill(tagID tagdm.TagID, evaluation uint8, years uint8) (*Skill, error) {
	skillId, err := NewSkillID()
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
	if err != nil {
		return nil, err
	}

	skillUpdatedAt := sharedvo.NewUpdatedAt()
	if err != nil {
		return nil, err
	}
	return &Skill{
		id:         skillId,
		tagID:      tagID,
		evaluation: eval,
		years:      y,
		createdAt:  skillCreatedAt,
		updatedAt:  skillUpdatedAt,
	}, nil
}

func ReconstructSkill(id string, tagID string, evaluation uint8, years uint8, createdAt time.Time, updatedAt time.Time) (*Skill, error) {
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

func (s *Skill) ID() SkillID {
	return s.id
}
func (s *Skill) TagID() tagdm.TagID {
	return s.tagID
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
func (s *Skill) Equal(other *Skill) bool {
	if other == nil {
		return false
	}

	return s.id.Equal(other.id) &&
		s.tagID.Equal(other.tagID) &&
		s.evaluation == other.evaluation &&
		s.years == other.years &&
		s.createdAt.Equal(other.createdAt) &&
		s.updatedAt.Equal(other.updatedAt)
}
func (s *Skill) SetID(id SkillID) {
	s.id = id
}

func (s *Skill) SetTagID(tagID tagdm.TagID) {
	s.tagID = tagID
}

func (s *Skill) SetEvaluation(evaluation uint8) error {
	eval, err := NewSkillEvaluation(evaluation)
	if err != nil {
		return err
	}
	s.evaluation = eval
	return nil
}

func (s *Skill) SetYears(years uint8) error {
	y, err := NewSkillYear(years)
	if err != nil {
		return err
	}
	s.years = y
	return nil
}

func (s *Skill) SetCreatedAt(createdAt time.Time) error {
	created, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return err
	}
	s.createdAt = created
	return nil
}

func (s *Skill) SetUpdatedAt(updatedAt time.Time) error {
	updated, err := sharedvo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		return err
	}
	s.updatedAt = updated
	return nil
}
