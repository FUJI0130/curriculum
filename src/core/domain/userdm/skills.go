package userdm

import (
	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
)

type Skills struct {
	id         SkillID            `db:"id"`
	tagId      tagdm.TagID        `db:"tag_id"`
	userId     UserID             `db:"user_id"`
	evaluation SkillsEvaluation   `db:"evaluation"`
	years      SkillsYears        `db:"years"`
	createdAt  sharedvo.CreatedAt `db:"created_at"`
	updatedAt  sharedvo.UpdatedAt `db:"updated_at"`
}

func NewSkills(tagId tagdm.TagID, userId UserID, evaluation uint8, years uint8) (*Skills, error) {
	eval, err := NewSkillsEvaluation(evaluation)
	if err != nil {
		return nil, err
	}

	y, err := NewSkillsYears(years)
	if err != nil {
		return nil, err
	}

	return &Skills{
		tagId:      tagId,
		userId:     userId,
		evaluation: *eval,
		years:      *y,
	}, nil
}
func (s *Skills) ID() SkillID {
	return s.id
}
func (s *Skills) TagID() tagdm.TagID {
	return s.tagId
}
func (s *Skills) UserID() UserID {
	return s.userId
}

func (s *Skills) Evaluation() SkillsEvaluation {
	return s.evaluation
}

func (s *Skills) Years() SkillsYears {
	return s.years
}
func (s *Skills) CreatedAt() sharedvo.CreatedAt {
	return s.createdAt
}

func (s *Skills) UpdatedAt() sharedvo.UpdatedAt {
	return s.updatedAt
}
