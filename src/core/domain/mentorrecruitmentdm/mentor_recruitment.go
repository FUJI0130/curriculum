package mentorrecruitmentdm

import (
	"errors"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
)

const descriptionMaxlength = 2000

type MentorRecruitment struct {
	id                  MentorRecruitmentID   `db:"id"`
	title               string                `db:"title"`
	categoryID          categorydm.CategoryID `db:"category_id"`
	budget              *Budget
	applicationPeriod   *ApplicationPeriod
	consultation_format int                `db:"consultation_format"`
	consultation_method int                `db:"consultation_method"`
	description         string             `db:"description"`
	status              int                `db:"status"`
	createdAt           sharedvo.CreatedAt `db:"created_at"`
	updatedAt           sharedvo.UpdatedAt `db:"updated_at"`
}

const (
	StatusOpen   = 0
	StatusClosed = 1
)

func NewMentorRecruitment(
	title string,
	categoryID categorydm.CategoryID,
	budgetFrom int,
	budgetTo int,
	applicationPeriodFrom time.Time,
	applicationPeriodTo time.Time,
	consultationFormat int,
	consultationMethod int,
	description string,
	status int,
) (*MentorRecruitment, error) {

	if title == "" {
		return nil, errors.New("title is empty")
	}

	if len(description) > descriptionMaxlength {
		return nil, errors.New("description is too long")
	}
	// Status のバリデーション
	if status != StatusOpen && status != StatusClosed {
		return nil, errors.New("invalid status: must be either open (0) or closed (1)")
	}
	budget, err := NewBudget(budgetFrom, budgetTo)
	if err != nil {
		return nil, err
	}

	applicationPeriod, err := NewApplicationPeriod(applicationPeriodFrom, applicationPeriodTo)
	if err != nil {
		return nil, err
	}

	mentorRecruitmentId, err := NewMentorRecruitmentID()
	if err != nil {
		return nil, err
	}

	createdAt := sharedvo.NewCreatedAt()
	updatedAt := sharedvo.NewUpdatedAt()

	return &MentorRecruitment{
		id:                  mentorRecruitmentId,
		title:               title,
		categoryID:          categoryID,
		budget:              budget,
		applicationPeriod:   applicationPeriod,
		consultation_format: consultationFormat,
		consultation_method: consultationMethod,
		description:         description,
		status:              status,
		createdAt:           createdAt,
		updatedAt:           updatedAt,
	}, nil
}

// Accessor methods for MentorRecruitments
func (mr *MentorRecruitment) ID() MentorRecruitmentID {
	return mr.id
}

func (mr *MentorRecruitment) Title() string {
	return mr.title
}

func (mr *MentorRecruitment) CategoryID() categorydm.CategoryID {
	return mr.categoryID
}

func (mr *MentorRecruitment) BudgetFrom() int {
	if mr.budget != nil {
		return mr.budget.From()
	}
	return 0 // または適切なデフォルト値、あるいはエラーを返す
}

func (mr *MentorRecruitment) BudgetTo() int {
	if mr.budget != nil {
		return mr.budget.To()
	}
	return 0 // または適切なデフォルト値、あるいはエラーを返す
}

func (mr *MentorRecruitment) ApplicationPeriodFrom() time.Time {
	if mr.applicationPeriod != nil {
		return mr.applicationPeriod.From()
	}
	return time.Time{} // または適切なデフォルト値、あるいはエラーを返す
}

func (mr *MentorRecruitment) ApplicationPeriodTo() time.Time {
	if mr.applicationPeriod != nil {
		return mr.applicationPeriod.To()
	}
	return time.Time{} // または適切なデフォルト値、あるいはエラーを返す
}

func (mr *MentorRecruitment) ConsultationFormat() int {
	return mr.consultation_format
}

func (mr *MentorRecruitment) ConsultationMethod() int {
	return mr.consultation_method
}

func (mr *MentorRecruitment) Description() string {
	return mr.description
}

func (mr *MentorRecruitment) Status() int {
	return mr.status
}

func (mr *MentorRecruitment) CreatedAt() sharedvo.CreatedAt {
	return mr.createdAt
}

func (mr *MentorRecruitment) UpdatedAt() sharedvo.UpdatedAt {
	return mr.updatedAt
}
