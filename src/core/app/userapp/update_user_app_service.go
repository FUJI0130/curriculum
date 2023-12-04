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

	var skillsToUpdate []userdm.Skill
	var careersToUpdate []userdm.Career
	updatedAt := time.Now() // 更新時刻の初期化

	// スキルの更新または新規作成
	for _, skillReq := range req.UpdateData.Skills {
		found := false
		for i, existingSkill := range user.Skills() {
			if existingSkill.ID().String() == skillReq.ID {
				// 既存のスキルを更新
				updatedSkill, err := userdm.GenSkillWhenUpdate(skillReq.ID, skillReq.TagID, skillReq.Evaluation, skillReq.Years, existingSkill.CreatedAt().DateTime())
				if err != nil {
					return err
				}
				user.Skills()[i] = *updatedSkill
				skillsToUpdate = append(skillsToUpdate, *updatedSkill)
				found = true
				break
			}
		}
		if !found {
			// 新規スキルの作成
			newSkill, err := userdm.GenSkillWhenCreate(skillReq.TagID, user.ID().String(), skillReq.Evaluation, skillReq.Years)
			if err != nil {
				return err
			}
			skillsToUpdate = append(skillsToUpdate, *newSkill)
		}
	}

	// キャリアの更新または新規作成
	for _, careerReq := range req.UpdateData.Careers {
		found := false
		for i, existingCareer := range user.Careers() {
			if existingCareer.ID().String() == careerReq.ID {
				// 既存のキャリアを更新
				updatedCareer, err := userdm.GenCareerWhenUpdate(careerReq.ID, careerReq.Detail, careerReq.AdFrom, careerReq.AdTo, existingCareer.CreatedAt().DateTime())
				if err != nil {
					return err
				}
				user.Careers()[i] = *updatedCareer
				careersToUpdate = append(careersToUpdate, *updatedCareer)
				found = true
				break
			}
		}
		if !found {
			// 新規キャリアの作成
			newCareer, err := userdm.GenCareerWhenCreate(careerReq.Detail, careerReq.AdFrom, careerReq.AdTo, user.ID().String())
			if err != nil {
				return err
			}
			careersToUpdate = append(careersToUpdate, *newCareer)
		}
	}

	// ユーザーエンティティ全体の更新
	if len(skillsToUpdate) > 0 || len(careersToUpdate) > 0 {
		updatedAt = time.Now()
	}
	if err := user.Update(req.UpdateData.Users.Name, req.UpdateData.Users.Email, req.UpdateData.Users.Password, req.UpdateData.Users.Profile, skillsToUpdate, careersToUpdate, updatedAt); err != nil {
		return err
	}

	// データベースに更新を反映
	return app.userRepo.Update(ctx, user)
}

func toUTCAndTruncate(t time.Time) time.Time {
	return t.UTC().Truncate(time.Second)
}
