package mentorrecruitmentdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

// ReconstructMentorRecruitment は既存のデータを元に MentorRecruitment エンティティを再構築します。
func GenWhenCreate(
	id MentorRecruitmentID,
	title string,
	categoryID categorydm.CategoryID,
	budgetFrom,
	budgetTo uint32,
	applicationPeriodFrom,
	applicationPeriodTo time.Time,
	consultationFormat,
	consultationMethod uint8,
	description string,
	status uint8,
	createdAt, updatedAt time.Time) (*MentorRecruitment, error) {
	if title == "" {
		return nil, customerrors.NewUnprocessableEntityError("title is empty")
	}
	budget, err := NewBudget(budgetFrom, budgetTo)
	if err != nil {
		return nil, err
	}

	applicationPeriod, err := NewApplicationPeriod(applicationPeriodFrom, applicationPeriodTo)
	if err != nil {
		return nil, err
	}

	if len(description) > descriptionMaxlength {
		return nil, customerrors.NewUnprocessableEntityError("description is too long")
	}
	if status != StatusOpen && status != StatusClosed {
		return nil, customerrors.NewUnprocessableEntityError("invalid status")
	}

	createdAtVal, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityError(err, "createdAt is invalid")
	}
	updatedAtVal, err := sharedvo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityError(err, "updatedAt is invalid")
	}

	return &MentorRecruitment{
		id:                  id,
		title:               title,
		categoryID:          categoryID,
		budget:              budget,
		applicationPeriod:   applicationPeriod,
		consultation_format: consultationFormat,
		consultation_method: consultationMethod,
		description:         description,
		status:              status,
		createdAt:           createdAtVal,
		updatedAt:           updatedAtVal,
	}, nil
}
