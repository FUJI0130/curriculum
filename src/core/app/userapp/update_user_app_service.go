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
	for _, skillReq := range req.UpdateData.Skills {
		tagID, ok := tagIDMap[skillReq.TagName]
		if !ok {
			return customerrors.NewNotFoundErrorf("Tag not found for TagName: %s", skillReq.TagName)
		}

		updated := false
		for i, skill := range user.Skills() {
			if skill.TagID().String() == tagID {
				if err := user.UpdateSkill(i, skill.ID(), tagdm.TagID(tagID), skillReq.Evaluation, skillReq.Years, skill.CreatedAt().DateTime(), time.Now()); err != nil {
					return err
				}
				updated = true
				break
			}
		}
		if !updated {
			newSkill, err := userdm.NewSkill(tagdm.TagID(tagID), skillReq.Evaluation, skillReq.Years)
			if err != nil {
				return err
			}
			// user.Skills = append(user.Skills, *newSkill)
			user.AppendSkill(*newSkill)
		}
	}

	// キャリアの更新
	for _, careerReq := range req.UpdateData.Careers {
		updated := false
		for i, career := range user.Careers() {
			if career.ID().String() == careerReq.ID {
				if err := user.UpdateCareer(i, career.ID(), careerReq.Detail, careerReq.AdFrom, careerReq.AdTo, career.CreatedAt().DateTime(), time.Now()); err != nil {
					return err
				}
				updated = true
				break
			}
		}
		if !updated {
			newCareer, err := userdm.NewCareer(careerReq.Detail, careerReq.AdFrom, careerReq.AdTo)
			if err != nil {
				return err
			}
			// user.Careers = append(user.Careers, *newCareer)
			user.AppendCareer(*newCareer)
		}
	}

	// ユーザーエンティティの更新
	if err := user.Update(req.UpdateData.Users.Name, req.UpdateData.Users.Email, req.UpdateData.Users.Password, req.UpdateData.Users.Profile, time.Now()); err != nil {
		return err
	}

	// データベースに更新を反映
	return app.userRepo.Update(ctx, user)
}

func toUTCAndTruncate(t time.Time) time.Time {
	return t.UTC().Truncate(time.Second)
}
