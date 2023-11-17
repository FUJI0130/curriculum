package rdbimpl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FUJI0130/curriculum/src/core/domain/mentorrecruitmentdm"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type mentorRecruitmentsTagsRepositoryImpl struct{}

func NewMentorRecruitmentsTagsRepository() mentorrecruitmentdm.MentorRecruitmentsTagsRepository {
	return &mentorRecruitmentsTagsRepositoryImpl{}
}

func (repo *mentorRecruitmentsTagsRepositoryImpl) Store(ctx context.Context, mentorRecruitmentID mentorrecruitmentdm.MentorRecruitmentID, tagID tagdm.TagID) error {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return errors.New("no transaction found in context")
	}
	query := "INSERT INTO mentor_recruitments_tags (mentor_recruitment_id, tag_id) VALUES (?, ?)"
	_, err := conn.Exec(query, mentorRecruitmentID.String(), tagID.String())

	if err != nil {
		return customerrors.WrapInternalServerError(err, "Failed to store mentor recruitment tag")
	}

	return nil
}

func (repo *mentorRecruitmentsTagsRepositoryImpl) FindByMentorRecruitmentID(ctx context.Context, mentorRecruitmentID mentorrecruitmentdm.MentorRecruitmentID) ([]tagdm.TagID, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return nil, errors.New("no transaction found in context")
	}

	query := "SELECT tag_id FROM mentor_recruitments_tags WHERE mentor_recruitment_id = ?"
	var tagIDs []string // 一時的にtag_idを格納するためのスライス

	err := conn.Select(&tagIDs, query, mentorRecruitmentID.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.WrapNotFoundError(err, "Mentor recruitment tags not found")
		}
		return nil, customerrors.WrapInternalServerError(err, "Error querying mentor recruitment tags")
	}

	// string型のtag_idをTagID型に変換
	var result []tagdm.TagID
	for _, idStr := range tagIDs {
		tagID, err := tagdm.NewTagIDFromString(idStr)
		if err != nil {
			return nil, customerrors.WrapUnprocessableEntityError(err, "Error converting tag ID")
		}
		result = append(result, tagID)
	}

	return result, nil
}
