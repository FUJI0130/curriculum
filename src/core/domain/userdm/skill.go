package userdm

import (
	"fmt"
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

func NewSkill(tagID tagdm.TagID, userID UserID, evaluation uint8, years uint8, createdAt time.Time, updatedAt time.Time) (*Skill, error) {
	eval, err := NewSkillEvaluation(evaluation)
	if err != nil {
		return nil, err
	}

	y, err := NewSkillYear(years)
	if err != nil {
		return nil, err
	}
	skillCreatedAt, err := sharedvo.NewCreatedAt(createdAt)
	if err != nil {
		return nil, err
	}

	skillUpdatedAt, err := sharedvo.NewUpdatedAt(updatedAt)
	if err != nil {
		fmt.Printf("Skill NewUpdatedAt  Time taken for updatedAt.Before(time.Now()): %v\n", sharedvo.LastDuration)
		return nil, err
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
