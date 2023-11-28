package rdbimpl

import (
	"context"
	"errors"
	"log"

	"github.com/FUJI0130/curriculum/src/core/domain/mentorrecruitmentdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type mentorRecruitmentCommandRepositoryImpl struct{}

func NewMentorRecruitmentCommandRepository() mentorrecruitmentdm.MentorRecruitmentCommandRepository {
	return &mentorRecruitmentCommandRepositoryImpl{}
}

func (repo *mentorRecruitmentCommandRepositoryImpl) Store(ctx context.Context, mentorRecruitment *mentorrecruitmentdm.MentorRecruitment) error {
	log.Printf("store start")
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return errors.New("no transaction found in context")
	}

	query := `INSERT INTO mentor_recruitments (
		id, 
		title, 
		category_id, 
		budget_from, 
		budget_to, 
		application_period_from, 
		application_period_to, 
		consultation_format, 
		consultation_method, 
		description, 
		status, 
		created_at, 
		updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := conn.Exec(
		query,
		mentorRecruitment.ID().String(),
		mentorRecruitment.Title(),
		mentorRecruitment.CategoryID().String(),
		mentorRecruitment.BudgetFrom(),
		mentorRecruitment.BudgetTo(),
		mentorRecruitment.ApplicationPeriodFrom(),
		mentorRecruitment.ApplicationPeriodTo(),
		mentorRecruitment.ConsultationFormat(),
		mentorRecruitment.ConsultationMethod(),
		mentorRecruitment.Description(),
		mentorRecruitment.Status(),
		mentorRecruitment.CreatedAt().DateTime(),
		mentorRecruitment.UpdatedAt().DateTime(),
	)

	if err != nil {
		log.Printf("Error storing mentor recruitment: %v\n", err)
		return customerrors.WrapInternalServerError(err, "Failed to store mentor recruitment")
	}

	return nil
}
