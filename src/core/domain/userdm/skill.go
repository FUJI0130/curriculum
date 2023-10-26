package userdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
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

func NewSkill(tagID tagdm.TagID, userID UserID, evaluation uint8, years uint8) (*Skill, error) {
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
		userID:     userID,
		evaluation: eval,
		years:      y,
		createdAt:  skillCreatedAt,
		updatedAt:  skillUpdatedAt,
	}, nil
}

func ReconstructSkill(id string, tagID string, userID string, evaluation uint8, years uint8, createdAt time.Time, updatedAt time.Time) (*Skill, error) {
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
func (s *Skill) Equal(other *Skill) bool {
	if other == nil {
		return false
	}

	return s.id.Equal(other.id) &&
		s.tagID.Equal(other.tagID) &&
		s.userID.Equal(other.userID) &&
		s.evaluation == other.evaluation &&
		s.years == other.years &&
		s.createdAt.Equal(other.createdAt) &&
		s.updatedAt.Equal(other.updatedAt)
}

func (s *Skill) MismatchedFields(other *Skill) map[string]bool {
	if other == nil {
		return map[string]bool{"other": true}
	}

	mismatches := map[string]bool{}

	if !s.id.Equal(other.id) {
		mismatches["id"] = true
	}
	if !s.tagID.Equal(other.tagID) {
		mismatches["tagID"] = true
	}
	if !s.userID.Equal(other.userID) {
		mismatches["userID"] = true
	}
	if s.evaluation != other.evaluation {
		mismatches["evaluation"] = true
	}
	if s.years != other.years {
		mismatches["years"] = true
	}
	if !s.createdAt.Equal(other.createdAt) {
		mismatches["createdAt"] = true
	}
	if !s.updatedAt.Equal(other.updatedAt) {
		mismatches["updatedAt"] = true
	}

	return mismatches
}
