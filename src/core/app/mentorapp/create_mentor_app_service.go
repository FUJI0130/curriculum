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
	tagDomainService         tagdm.TagDomainService
}

func NewCreateMentorRecruitmentAppService(
	mentorRecruitmentRepo mentorrecruitmentdm.MentorRecruitmentRepository,
	mentorRecruitmentTagRepo mentorrecruitmentdm.MentorRecruitmentsTagsRepository,
	categoryRepo categorydm.CategoryRepository,
	tagDomainService tagdm.TagDomainService,
) *CreateMentorRecruitmentAppService {
	return &CreateMentorRecruitmentAppService{
		mentorRecruitmentRepo:    mentorRecruitmentRepo,
		mentorRecruitmentTagRepo: mentorRecruitmentTagRepo,
		categoryRepo:             categoryRepo,
		tagDomainService:         tagDomainService,
	}
}

func (app *CreateMentorRecruitmentAppService) Exec(ctx context.Context, req *CreateMentorRecruitmentRequest) (err error) {

	category, err := app.categoryRepo.FindByID(ctx, req.CategoryID)
	if err != nil {
		if customerrors.IsNotFoundError(err) {
			return err
		}
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
	for i, tagID := range req.TagIDs {
		tagName := ""
		if len(req.TagNames) > i {
			tagName = req.TagNames[i]
		}

		tag, err := app.tagDomainService.ProcessTag(ctx, tagID, tagName)
		if err != nil {
			return err
		}

		mentorRecruitmentTag := mentorrecruitmentdm.NewMentorRecruitmentTag(mentorRecruitment.ID(), tag.ID(), time.Now())
		if err = app.mentorRecruitmentTagRepo.Store(ctx, mentorRecruitmentTag); err != nil {
			return customerrors.WrapInternalServerError(err, "Failed to store mentor recruitment tag")
		}
	}
	return nil
}
