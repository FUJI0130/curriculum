package userapp

import (
	"context"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type UpdateUserAppService struct {
	userRepo userdm.UserRepository
	tagRepo  tagdm.TagRepository
}

func NewUpdateUserAppService(userRepo userdm.UserRepository, tagRepo tagdm.TagRepository) *UpdateUserAppService {
	return &UpdateUserAppService{
		userRepo: userRepo,
		tagRepo:  tagRepo,
	}
}

func (app *UpdateUserAppService) ExecUpdate(ctx context.Context, req *UpdateUserRequestData) error {
	user, err := app.userRepo.FindByUserID(ctx, req.UpdateData.Users.ID)
	if err != nil {
		return customerrors.NewNotFoundErrorf("User not found by UserID: %s", req.UpdateData.Users.ID)
	}

	// タグ名からtagIDを取得
	tagNames := make([]string, len(req.UpdateData.Skills))
	for i, skillReq := range req.UpdateData.Skills {
		tagNames[i] = skillReq.TagName
	}
	tags, err := app.tagRepo.FindByNames(ctx, tagNames)
	if err != nil {
		return err
	}

	// タグ名からtagIDのマップを作成
	tagIDMap := make(map[string]string)
	for _, tag := range tags {
		tagIDMap[tag.Name()] = tag.ID().String()
	}

	// スキルの更新
	var updatedSkills []*userdm.Skill
	for _, skillReq := range req.UpdateData.Skills {
		tagID, ok := tagIDMap[skillReq.TagName]
		if !ok {
			return customerrors.NewNotFoundErrorf("Tag not found for TagName: %s", skillReq.TagName)
		}

		// 既存のスキルを探して更新、または新しいスキルを作成
		found := false
		for _, existingSkill := range user.Skills() {
			if existingSkill.TagID().String() == tagID {
				updatedSkill, err := userdm.ReconstructSkill(existingSkill.ID().String(), tagID, skillReq.Evaluation, skillReq.Years, existingSkill.CreatedAt().DateTime(), time.Now())
				if err != nil {
					return err
				}
				updatedSkills = append(updatedSkills, updatedSkill)
				found = true
				break
			}
		}
		if !found {
			// 新しいスキルの作成
			newSkill, err := userdm.NewSkill(tagdm.TagID(tagID), skillReq.Evaluation, skillReq.Years)
			if err != nil {
				return err
			}
			updatedSkills = append(updatedSkills, newSkill)
		}
	}

	// キャリアの更新
	var updatedCareers []*userdm.Career
	for _, careerReq := range req.UpdateData.Careers {
		career, err := userdm.ReconstructCareer(careerReq.ID, careerReq.Detail, careerReq.AdFrom, careerReq.AdTo, time.Now(), time.Now())
		if err != nil {
			return err
		}
		updatedCareers = append(updatedCareers, career)
	}

	// ユーザーエンティティの更新
	if err := user.Update(req.UpdateData.Users.Name, req.UpdateData.Users.Email, req.UpdateData.Users.Password, req.UpdateData.Users.Profile, updatedSkills, updatedCareers, time.Now()); err != nil {
		return err
	}

	// データベースに更新を反映
	return app.userRepo.Update(ctx, user)
}

func toUTCAndTruncate(t time.Time) time.Time {
	return t.UTC().Truncate(time.Second)
}
