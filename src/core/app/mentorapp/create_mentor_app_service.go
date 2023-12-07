package mentorapp

import (
	"context"
	"database/sql"
	"time"

	"github.com/cockroachdb/errors"

	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
	"github.com/FUJI0130/curriculum/src/core/domain/mentorrecruitmentdm"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type CreateMentorRecruitmentAppService struct {
	mentorRecruitmentRepo    mentorrecruitmentdm.MentorRecruitmentRepository
	mentorRecruitmentTagRepo mentorrecruitmentdm.MentorRecruitmentsTagsRepository
	tagRepo                  tagdm.TagRepository
	categoryRepo             categorydm.CategoryRepository
}

func NewCreateMentorRecruitmentAppService(
	mentorRecruitmentRepo mentorrecruitmentdm.MentorRecruitmentRepository,
	mentorRecruitmentTagRepo mentorrecruitmentdm.MentorRecruitmentsTagsRepository,
	tagRepo tagdm.TagRepository,
	categoryRepo categorydm.CategoryRepository,
) *CreateMentorRecruitmentAppService {
	return &CreateMentorRecruitmentAppService{
		mentorRecruitmentRepo:    mentorRecruitmentRepo,
		mentorRecruitmentTagRepo: mentorRecruitmentTagRepo,
		tagRepo:                  tagRepo,
		categoryRepo:             categoryRepo,
	}
}

func (app *CreateMentorRecruitmentAppService) Exec(ctx context.Context, req *CreateMentorRecruitmentRequest) (err error) {

	category, err := app.categoryRepo.FindByName(ctx, req.CategoryName)

	if errors.Is(err, sql.ErrNoRows) {
		newCategory, err := categorydm.GenWhenCreate(req.CategoryName)
		if err != nil {
			return customerrors.WrapInternalServerError(err, "新しいカテゴリの作成に失敗しました")
		}

		if err = app.categoryRepo.Store(ctx, newCategory); err != nil {
			return err
		}
		category = newCategory
	} else if err != nil {
		return err
	}

	tagIDs := make([]string, 0, len(req.TagNames))
	for _, tagName := range req.TagNames {
		tag, err := app.tagRepo.FindByName(ctx, tagName)
		if err != nil {
			return err
		}
		if tag == nil {
			newTag, err := tagdm.GenWhenCreateTag(tagName)
			if err != nil {
				return err
			}
			//N+1
			if err = app.tagRepo.Store(ctx, newTag); err != nil {
				return err
			}
			tagIDs = append(tagIDs, newTag.ID().String())
		} else {
			tagIDs = append(tagIDs, tag.ID().String())
		}
	}

	mentorRecruitment, err := mentorrecruitmentdm.NewMentorRecruitment(
		req.Title,
		category.ID(),
		req.BudgetFrom,
		req.BudgetTo,
		req.ApplicationPeriodFrom,
		req.ApplicationPeriodTo,
		req.ConsultationFormat,
		req.ConsultationMethod,
		req.Description,
		req.Status,
	)
	if err != nil {
		return err
	}

	err = app.mentorRecruitmentRepo.Store(ctx, mentorRecruitment)
	if err != nil {
		return err
	}

	for _, tagIDStr := range tagIDs {
		tagID, err := tagdm.NewTagIDFromString(tagIDStr)
		if err != nil {
			return customerrors.WrapUnprocessableEntityError(err, "Invalid tag ID format")
		}

		mentorRecruitmentTag := mentorrecruitmentdm.NewMentorRecruitmentTag(mentorRecruitment.ID(), tagID, time.Now())
		if err = app.mentorRecruitmentTagRepo.Store(ctx, mentorRecruitmentTag); err != nil {
			return customerrors.WrapInternalServerError(err, "Failed to store mentor recruitment tag")
		}
	}

	return nil
}
