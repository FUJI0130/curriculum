package rdbimpl

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/mentorrecruitmentdm"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type mentorRecruitmentsTagsRepositoryImpl struct{}

func NewMentorRecruitmentsTagsRepository() mentorrecruitmentdm.MentorRecruitmentsTagsRepository {
	return &mentorRecruitmentsTagsRepositoryImpl{}
}

func (repo *mentorRecruitmentsTagsRepositoryImpl) Store(ctx context.Context, mentorRecruitmentTag *mentorrecruitmentdm.MentorRecruitmentTag) error {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return errors.New("no transaction found in context")
	}

	query := "INSERT INTO mentor_recruitments_tags (mentor_recruitment_id, tag_id, created_at) VALUES (?, ?, ?)"
	_, err := conn.Exec(
		query,
		mentorRecruitmentTag.MentorRecruitmentID.String(),
		mentorRecruitmentTag.TagID.String(),
		mentorRecruitmentTag.CreatedAt,
	)

	if err != nil {
		return customerrors.WrapInternalServerError(err, "Failed to store mentor recruitment tag")
	}

	return nil
}

func (repo *mentorRecruitmentsTagsRepositoryImpl) FindByMentorRecruitmentID(ctx context.Context, mentorRecruitmentID mentorrecruitmentdm.MentorRecruitmentID) ([]*mentorrecruitmentdm.MentorRecruitmentTag, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return nil, errors.New("no transaction found in context")
	}

	query := "SELECT mentor_recruitment_id, tag_id, created_at FROM mentor_recruitments_tags WHERE mentor_recruitment_id = ?"
	var rows []struct {
		MentorRecruitmentID string
		TagID               string
		CreatedAt           time.Time
	}

	err := conn.Select(&rows, query, mentorRecruitmentID.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.WrapNotFoundError(err, "Mentor recruitment tags not found")
		}
		return nil, customerrors.WrapInternalServerError(err, "Error querying mentor recruitment tags")
	}

	var result []*mentorrecruitmentdm.MentorRecruitmentTag
	for _, row := range rows {
		mentorRecruitmentID, _ := mentorrecruitmentdm.NewMentorRecruitmentIDFromString(row.MentorRecruitmentID)
		tagID, _ := tagdm.NewTagIDFromString(row.TagID)

		result = append(result, mentorrecruitmentdm.NewMentorRecruitmentTag(mentorRecruitmentID, tagID, row.CreatedAt))
	}

	return result, nil
}
