package mentorapp

import (
	"context"
	"time"

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
	category, err := app.categoryRepo.FindByID(ctx, req.CategoryID)
	if err != nil {
		return customerrors.WrapInternalServerError(err, "カテゴリの検索に失敗しました")
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
		return customerrors.WrapInternalServerError(err, "メンター募集の新規作成に失敗しました")
	}

	err = app.mentorRecruitmentRepo.Store(ctx, mentorRecruitment)
	if err != nil {
		return customerrors.WrapInternalServerError(err, "メンター募集の保存に失敗しました")
	}

	for _, tagIDStr := range req.TagIDs {
		tagID, err := tagdm.NewTagIDFromString(tagIDStr)
		if err != nil {
			return customerrors.WrapUnprocessableEntityError(err, "無効なタグIDフォーマット")
		}

		_, err = app.tagRepo.FindByID(ctx, tagID.String())
		if err != nil {
			return customerrors.WrapInternalServerError(err, "タグの検索に失敗しました")
		}

		mentorRecruitmentTag := mentorrecruitmentdm.NewMentorRecruitmentTag(mentorRecruitment.ID(), tagID, time.Now())
		if err = app.mentorRecruitmentTagRepo.Store(ctx, mentorRecruitmentTag); err != nil {
			return customerrors.WrapInternalServerError(err, "メンター募集タグの保存に失敗しました")
		}
	}

	return nil
}
