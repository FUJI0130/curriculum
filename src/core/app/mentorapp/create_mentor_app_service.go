package mentorapp

import (
	"context"

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

func NewCreateMentorRecruitmentAppService(mentorRecruitmentRepo mentorrecruitmentdm.MentorRecruitmentRepository, tagRepo tagdm.TagRepository, categoryRepo categorydm.CategoryRepository) *CreateMentorRecruitmentAppService {
	return &CreateMentorRecruitmentAppService{
		mentorRecruitmentRepo: mentorRecruitmentRepo,
		tagRepo:               tagRepo,
		categoryRepo:          categoryRepo,
	}
}

func (app *CreateMentorRecruitmentAppService) Exec(ctx context.Context, req *CreateMentorRecruitmentRequest) (err error) {
	// カテゴリの確認
	category, err := app.categoryRepo.FindByID(ctx, req.CategoryID)
	if err != nil {
		return err
	}
	if category == nil {
		return customerrors.NewUnprocessableEntityError("category not found")
	}

	// タグの処理
	tagIds := make([]string, 0, len(req.TagNames))
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
			if err = app.tagRepo.Store(ctx, newTag); err != nil {
				return err
			}
			tagIds = append(tagIds, newTag.ID().String())
		} else {
			tagIds = append(tagIds, tag.ID().String())
		}
	}

	// メンター募集の作成
	mentorRecruitment, err := mentorrecruitmentdm.NewMentorRecruitment(
		// メンター募集の詳細情報
		req.Title,
		categorydm.CategoryID(req.CategoryID),
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

	// メンター募集の保存
	err = app.mentorRecruitmentRepo.Store(ctx, mentorRecruitment)
	if err != nil {
		return err
	}

	// 中間テーブルの更新
	for _, tagID := range tagIds {
		makeTag, err := tagdm.GenWhenCreateTag(tagID)
		err = app.mentorRecruitmentTagRepo.Store(ctx, mentorRecruitment.ID(), tagdm.TagID(makeTag.ID().String()))
		if err != nil {
			return err
		}
	}

	return nil
}
