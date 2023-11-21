package rdbimpl

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
	"github.com/FUJI0130/curriculum/src/core/domain/mentorrecruitmentdm"
	"github.com/FUJI0130/curriculum/src/core/infra/datamodel"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type mentorRecruitmentRepositoryImpl struct{}

func NewMentorRecruitmentRepository() mentorrecruitmentdm.MentorRecruitmentRepository {
	return &mentorRecruitmentRepositoryImpl{}
}

func (repo *mentorRecruitmentRepositoryImpl) Store(ctx context.Context, mentorRecruitment *mentorrecruitmentdm.MentorRecruitment) error {
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

func (repo *mentorRecruitmentRepositoryImpl) FindByID(ctx context.Context, id mentorrecruitmentdm.MentorRecruitmentID) (*mentorrecruitmentdm.MentorRecruitment, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return nil, errors.New("no transaction found in context")
	}

	query := "SELECT * FROM mentor_recruitments WHERE id = ?"
	var tempMentorRecruitment datamodel.MentorRecruitment
	err := conn.Get(&tempMentorRecruitment, query, id.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.WrapNotFoundError(err, "Mentor recruitment not found")
		}
		return nil, err
	}

	mentorRecruitment, err := mentorrecruitmentdm.GenWhenCreate(
		mentorrecruitmentdm.MentorRecruitmentID(tempMentorRecruitment.ID),
		tempMentorRecruitment.Title,
		categorydm.CategoryID(tempMentorRecruitment.CategoryID),
		tempMentorRecruitment.BudgetFrom,
		tempMentorRecruitment.BudgetTo,
		tempMentorRecruitment.ApplicationPeriodFrom,
		tempMentorRecruitment.ApplicationPeriodTo,
		tempMentorRecruitment.ConsultationFormat,
		tempMentorRecruitment.ConsultationMethod,
		tempMentorRecruitment.Description,
		tempMentorRecruitment.Status,
		tempMentorRecruitment.CreatedAt,
		tempMentorRecruitment.UpdatedAt,
	)
	if err != nil {
		return nil, customerrors.WrapInternalServerError(err, "Error reconstructing mentor recruitment")
	}

	return mentorRecruitment, nil
}

func (repo *mentorRecruitmentRepositoryImpl) FindByCategoryID(ctx context.Context, categoryID categorydm.CategoryID) ([]*mentorrecruitmentdm.MentorRecruitment, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return nil, errors.New("no transaction found in context")
	}

	query := "SELECT * FROM mentor_recruitments WHERE category_id = ?"
	var tempMentorRecruitments []datamodel.MentorRecruitment
	err := conn.Select(&tempMentorRecruitments, query, categoryID.String())
	if err != nil {
		return nil, customerrors.WrapInternalServerError(err, "Error querying mentor recruitments by category")
	}

	var mentorRecruitments []*mentorrecruitmentdm.MentorRecruitment
	for _, temp := range tempMentorRecruitments {
		mentorRecruitment, err := mentorrecruitmentdm.GenWhenCreate(
			mentorrecruitmentdm.MentorRecruitmentID(temp.ID),
			temp.Title,
			categorydm.CategoryID(temp.CategoryID),
			temp.BudgetFrom,
			temp.BudgetTo,
			temp.ApplicationPeriodFrom,
			temp.ApplicationPeriodTo,
			temp.ConsultationFormat,
			temp.ConsultationMethod,
			temp.Description,
			temp.Status,
			temp.CreatedAt,
			temp.UpdatedAt,
		)
		if err != nil {
			return nil, customerrors.WrapInternalServerError(err, "Error generating mentor recruitment")
		}
		mentorRecruitments = append(mentorRecruitments, mentorRecruitment)
	}

	return mentorRecruitments, nil
}
