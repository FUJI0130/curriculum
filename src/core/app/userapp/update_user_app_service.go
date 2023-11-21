package userapp

import (
	"context"
	"log"
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
	log.Printf("req is : %v\n", req)

	userDomainDataOnDB, err := app.userRepo.FindByUserID(ctx, req.UpdateData.Users.ID)
	if err != nil {
		return customerrors.NewNotFoundErrorf("Update_user_app_service  Exec User Not Exist  UserID is : %s", req.UpdateData.Users.ID)
	}
	log.Printf("log checkpoint1")

	updatedUser, err := app.updateUsers(ctx, userDomainDataOnDB, &req.UpdateData.Users)
	if err != nil {
		return err
	}
	log.Print("log checkpoint2")

	updatedSkills, err := app.updateSkills(ctx, userDomainDataOnDB, req.UpdateData.Skills)
	if err != nil {
		return err
	}
	log.Print("log checkpoint3")

	updatedCareers, err := app.updateCareers(ctx, userDomainDataOnDB, req.UpdateData.Careers)
	if err != nil {
		return err
	}

	user, err := userdm.GenWhenUpdate(updatedUser, updatedSkills, updatedCareers)
	if err != nil {
		return err
	}

	return app.userRepo.Update(ctx, user)
}

func (app *UpdateUserAppService) updateUsers(ctx context.Context, userDataOnDB *userdm.User, req *UpdateUserRequest) (*userdm.User, error) {
	updatedUser, err := userdm.GenUserWhenUpdate(
		userDataOnDB.ID().String(),
		req.Name,
		req.Email,
		req.Password,
		req.Profile,
		userDataOnDB.CreatedAt().DateTime(),
	)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (app *UpdateUserAppService) updateSkills(ctx context.Context, user *userdm.User, skillsReq []UpdateSkillRequest) ([]*userdm.Skill, error) {
	// 既存のスキルのマップを作成
	existingSkillsMap := make(map[string]userdm.Skill)
	for _, skill := range user.Skills() {
		existingSkillsMap[skill.ID().String()] = skill
	}

	// タグ名のリストとタグ情報の取得
	tagNames := make([]string, 0, len(skillsReq))
	for _, s := range skillsReq {
		tagNames = append(tagNames, s.TagName)
	}
	tags, err := app.tagRepo.FindByNames(ctx, tagNames)
	if err != nil {
		return nil, err
	}
	tagsMap := make(map[string]tagdm.TagID)
	for _, tag := range tags {
		tagsMap[tag.Name()] = tag.ID()
	}

	// 更新するスキルのリストを生成
	var updatedSkills []*userdm.Skill
	for _, s := range skillsReq {
		tagID, exists := tagsMap[s.TagName]
		if !exists {
			return nil, customerrors.NewUnprocessableEntityErrorf("No tag found with name: %s", s.TagName)
		}

		existingSkill, exists := existingSkillsMap[s.ID]
		if !exists {
			return nil, customerrors.NewNotFoundErrorf("Skill not found with ID: %s", s.ID)
		}

		updatedSkill, err := userdm.GenSkillWhenUpdate(
			s.ID,
			tagID.String(),
			user.ID().String(),
			s.Evaluation,
			s.Years,
			existingSkill.CreatedAt().DateTime(),
			time.Now(),
		)
		if err != nil {
			return nil, err
		}
		updatedSkills = append(updatedSkills, updatedSkill)
	}

	return updatedSkills, nil
}

func (app *UpdateUserAppService) updateCareers(ctx context.Context, user *userdm.User, careersReq []UpdateCareerRequest) ([]*userdm.Career, error) {

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	existingCareers := user.Careers()

	existingCareersMap := make(map[string]*userdm.Career)
	for _, career := range existingCareers {
		careerCopy := career
		existingCareersMap[career.ID().String()] = &careerCopy
	}

	careersToUpdate := make([]*userdm.Career, 0, len(careersReq))

	for _, c := range careersReq {
		adFromInJST := c.AdFrom.In(jst)
		adToInJST := c.AdTo.In(jst)
		var career *userdm.Career

		if existingCareer, ok := existingCareersMap[c.ID]; ok {
			career, err = userdm.GenCareerWhenUpdate(
				existingCareer.ID().String(),
				c.Detail,
				adFromInJST,
				adToInJST,
				user.ID().String(),
				existingCareer.CreatedAt().DateTime(),
				time.Now(),
			)
			if err != nil {
				return nil, err
			}
		} else {
			career, err = userdm.GenCareerWhenCreate(
				c.Detail,
				adFromInJST,
				adToInJST,
				user.ID().String(),
				time.Now(),
				time.Now(),
			)
			if err != nil {
				return nil, err
			}
		}
		careersToUpdate = append(careersToUpdate, career)
	}

	return careersToUpdate, nil
}

func toUTCAndTruncate(t time.Time) time.Time {
	return t.UTC().Truncate(time.Second)
}
