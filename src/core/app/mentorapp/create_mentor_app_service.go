package mentorapp

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/cockroachdb/errors"

	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
	"github.com/FUJI0130/curriculum/src/core/domain/mentorrecruitmentdm"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type CreateMentorRecruitmentAppService struct {
	mentorRecruitmentRepo    mentorrecruitmentdm.MentorRecruitmentCommandRepository
	mentorRecruitmentTagRepo mentorrecruitmentdm.MentorRecruitmentsTagsRepository
	tagRepo                  tagdm.TagRepository
	categoryRepo             categorydm.CategoryRepository
}

func NewCreateMentorRecruitmentAppService(
	mentorRecruitmentRepo mentorrecruitmentdm.MentorRecruitmentCommandRepository,
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
	log.Println("Start Exec in CreateMentorRecruitmentAppService")

	// カテゴリの確認
	log.Println("Checking category:", req.CategoryName)
	category, err := app.categoryRepo.FindByName(ctx, req.CategoryName)

	// sql.ErrNoRows が返された場合、新しいカテゴリを作成
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("Category not found, creating new category:", req.CategoryName)
		newCategory, err := categorydm.GenWhenCreate(req.CategoryName)
		if err != nil {
			log.Println("Error creating new category:", err)
			return customerrors.WrapInternalServerError(err, "新しいカテゴリの作成に失敗しました")
		}

		// 新しいカテゴリを保存
		if err = app.categoryRepo.Store(ctx, newCategory); err != nil {
			log.Println("Error storing new category:", err)
			return err
		}
		category = newCategory
	} else if err != nil {
		// 他のエラーの場合
		log.Println("Error finding category:", err)
		return err
	}

	// タグの処理
	log.Println("Processing tags:", req.TagNames)
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

	log.Println("Creating mentor recruitment")
	mentorRecruitment, err := mentorrecruitmentdm.NewMentorRecruitment(
		// メンター募集の詳細情報
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

	// メンター募集の内容をログに出力
	log.Printf("Mentor Recruitment to store: %+v\n", mentorRecruitment)

	// メンター募集の保存
	err = app.mentorRecruitmentRepo.Store(ctx, mentorRecruitment)
	if err != nil {
		return err
	}

	// 中間テーブルの更新
	for _, tagIDStr := range tagIds {
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
