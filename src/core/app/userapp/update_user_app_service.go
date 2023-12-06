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

	// tagNames := make([]string, 0, len(req.Skills))
	skillsToUpdate := make([]userdm.Skill, 0, len(req.UpdateData.Skills))
	// var careersToUpdate []userdm.Career
	careersToUpdate := make([]userdm.Career, 0, len(req.UpdateData.Careers))

	// スキルの更新または新規作成
	for _, skillReq := range req.UpdateData.Skills {
		found := false
		for i, existingSkill := range user.Skills() {
			if existingSkill.ID().String() == skillReq.ID {
				// 既存のスキルを更新
				convSkillID, err := convReqToSkillID(&skillReq)
				if err != nil {
					return err
				}
				convTagID, err := convReqToTagID(&skillReq)
				if err != nil {
					return err
				}
				updatedSkill, err := userdm.GenSkillWhenUpdate(*convSkillID, *convTagID, skillReq.Evaluation, skillReq.Years, existingSkill.CreatedAt().DateTime())
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
				convCareerID, err := convReqToCareerID(&careerReq)
				if err != nil {
					return err
				}

				updatedCareer, err := userdm.GenCareerWhenUpdate(*convCareerID, careerReq.Detail, careerReq.AdFrom, careerReq.AdTo, existingCareer.CreatedAt().DateTime())
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
	if err := user.Update(req.UpdateData.Users.Name, req.UpdateData.Users.Email, req.UpdateData.Users.Password, req.UpdateData.Users.Profile, skillsToUpdate, careersToUpdate); err != nil {
		return err
	}

	// データベースに更新を反映
	return app.userRepo.Update(ctx, user)
}

func toUTCAndTruncate(t time.Time) time.Time {
	return t.UTC().Truncate(time.Second)
}

// リクエストからskillIDとtagIDのvalueobjectを生成するプライベート関数
func convReqToSkillID(req *UpdateSkillRequest) (*userdm.SkillID, error) {

	var skillID userdm.SkillID
	skillID, err := userdm.NewSkillIDFromString(req.ID)
	if err != nil {
		return nil, err
	}

	return &skillID, nil

}
func convReqToTagID(req *UpdateSkillRequest) (*tagdm.TagID, error) {
	var tagID tagdm.TagID
	tagID, err := tagdm.NewTagIDFromString(req.TagID)
	if err != nil {
		return nil, err
	}
	return &tagID, nil
}

// リクエストからcareerIDのvalueobjectを生成するプライベート関数
func convReqToCareerID(req *UpdateCareerRequest) (*userdm.CareerID, error) {
	var careerID userdm.CareerID
	careerID, err := userdm.NewCareerIDFromString(req.ID)
	if err != nil {
		return nil, err
	}
	return &careerID, nil
}
