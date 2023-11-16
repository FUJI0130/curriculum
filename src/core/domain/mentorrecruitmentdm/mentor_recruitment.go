package mentorrecruitmentdm

import (
	"errors"

	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
)

const descriptionMaxlength = 2000

type MentorRecruitment struct {
	id                    MentorRecruitmentID   `db:"id"`
	title                 string                `db:"name"`
	categoryID            categorydm.CategoryID `db:"tag_id"`
	budgetFrom            int                   `db:"budget_from"`
	budgetTo              int                   `db:"budget_to"`
	applicationPeriodFrom int                   `db:"application_period_from"`
	applicationPeriodTo   int                   `db:"application_period_to"`
	consultation_format   int                   `db:"consultation_format"`
	consultation_method   int                   `db:"consultation_method"`
	description           string                `db:"description"`
	status                int                   `db:"status"`
	profile               string                `db:"profile"`
	createdAt             sharedvo.CreatedAt    `db:"created_at"`
	updatedAt             sharedvo.UpdatedAt    `db:"updated_at"`
}

func NewMentorRecruitment(
	title string,
	categoryID categorydm.CategoryID,
	budgetFrom int,
	budgetTo int,
	applicationPeriodFrom int,
	applicationPeriodTo int,
	consultationFormat int,
	consultationMethod int,
	description string,
	profile string,
) (*MentorRecruitment, error) {

	if title == "" {
		return nil, errors.New("title is empty")
	}

	if len(description) > descriptionMaxlength {
		return nil, errors.New("description is too long")
	}

	if profile == "" {
		return nil, errors.New("profile is empty")
	}

	// ここで Budget と ApplicationPeriod のバリデーションを行う
	// 例: budgetFrom が budgetTo より大きい場合、エラーを返す
	if budgetFrom > budgetTo {
		return nil, errors.New("budgetFrom cannot be greater than budgetTo")
	}

	// applicationPeriod のバリデーション (具体的な条件はビジネスルールに依存)
	// 例: applicationPeriodFrom が applicationPeriodTo より大きい場合、エラーを返す
	if applicationPeriodFrom > applicationPeriodTo {
		return nil, errors.New("applicationPeriodFrom cannot be greater than applicationPeriodTo")
	}

	mentorRecruitmentId, err := NewMentorRecruitmentID()
	if err != nil {
		return nil, err
	}

	createdAt := sharedvo.NewCreatedAt()
	updatedAt := sharedvo.NewUpdatedAt()

	return &MentorRecruitment{
		id:                    mentorRecruitmentId,
		title:                 title,
		categoryID:            categoryID,
		budgetFrom:            budgetFrom,
		budgetTo:              budgetTo,
		applicationPeriodFrom: applicationPeriodFrom,
		applicationPeriodTo:   applicationPeriodTo,
		consultation_format:   consultationFormat,
		consultation_method:   consultationMethod,
		description:           description,
		profile:               profile,
		createdAt:             createdAt,
		updatedAt:             updatedAt,
	}, nil
}

// Accessor methods for MentorRecruitments
func (mr *MentorRecruitment) ID() MentorRecruitmentID {
	return mr.id
}

func (mr *MentorRecruitment) Title() string {
	return mr.title
}
