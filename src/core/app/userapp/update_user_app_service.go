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

type UpdateUserRequestData struct {
	UpdateData struct {
		Users   updateUserRequest     `json:"Users"`
		Skills  []updateSkillRequest  `json:"Skills"`
		Careers []updateCareerRequest `json:"Careers"`
	} `json:"updateData"`
}

func (app *UpdateUserAppService) ExecUpdate(ctx context.Context, req *UpdateUserRequestData) error {
	log.Printf("req is : %v\n", req)

	userDataOnDB, err := app.userRepo.FindByUserID(ctx, req.UpdateData.Users.ID)
	if err != nil {
		return customerrors.NewNotFoundErrorf("Update_user_app_service  Exec User Not Exist  UserID is : %s", req.UpdateData.Users.ID)
	}
	log.Printf("log checkpoint1")

	updatedUser, err := app.updateUsers(ctx, userDataOnDB, &req.UpdateData.Users)
	if err != nil {
		return err
	}
	log.Print("log checkpoint2")

	updatedSkills, err := app.updateSkills(ctx, updatedUser, req.UpdateData.Skills)
	if err != nil {
		return err
	}
	log.Print("log checkpoint3")

	updatedCareers, err := app.updateCareers(ctx, updatedUser.ID(), req.UpdateData.Careers)
	if err != nil {
		return err
	}

	userDomain, err := userdm.GenWhenUpdate(updatedUser, updatedSkills, updatedCareers)
	if err != nil {
		return err
	}

	return app.userRepo.Update(ctx, userDomain)
}

func (app *UpdateUserAppService) updateUsers(ctx context.Context, userDataOnDB *userdm.User, req *updateUserRequest) (*userdm.User, error) {
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

func (app *UpdateUserAppService) updateSkills(ctx context.Context, userDataOnDB *userdm.User, skillsReq []updateSkillRequest) ([]*userdm.Skill, error) {
	updatedSkills := []*userdm.Skill{}
	tagNames := make([]string, 0, len(skillsReq))
	seenSkills := make(map[string]bool)

	for _, s := range skillsReq {
		if seenSkills[s.TagName] {
			return nil, customerrors.NewUnprocessableEntityErrorf("Skill with tag name %s is duplicated", s.TagName)
		}
		seenSkills[s.TagName] = true
		tagNames = append(tagNames, s.TagName)
	}

	tags, err := app.tagRepo.FindByNames(ctx, tagNames)
	if err != nil {
		return nil, err
	}

	tagsMap := make(map[string]*tagdm.Tag)
	for _, tag := range tags {
		tagsMap[tag.Name()] = tag
	}

	for _, tagName := range tagNames {
		if _, exists := tagsMap[tagName]; !exists {
			tag, err := tagdm.GenWhenCreateTag(tagName)
			if err != nil {
				return nil, err
			}
			//N+1
			if err = app.tagRepo.Store(ctx, tag); err != nil {
				return nil, err
			}
			tagsMap[tagName] = tag
		}
	}

	skills, err := app.userRepo.FindSkillsByUserID(ctx, userDataOnDB.ID().String())
	if err != nil {
		return nil, err
	}

	existingSkillsMap := make(map[string]*userdm.Skill)
	for _, skill := range skills {
		skillCopy := skill
		existingSkillsMap[skill.TagID().String()] = &skillCopy
	}

	for _, s := range skillsReq {
		tagID := tagsMap[s.TagName].ID().String()
		if skill, exists := existingSkillsMap[tagID]; exists {
			updatedSkill, err := userdm.GenSkillWhenUpdate(skill.ID().String(), tagID, skill.UserID().String(), s.Evaluation, s.Years, skill.CreatedAt().DateTime(), time.Now())
			if err != nil {
				return nil, err
			}
			updatedSkills = append(updatedSkills, updatedSkill)
		} else {
			newSkill, err := userdm.GenSkillWhenCreate(tagID, userDataOnDB.ID().String(), s.Evaluation, s.Years, time.Now(), time.Now())
			if err != nil {
				return nil, err
			}
			updatedSkills = append(updatedSkills, newSkill)
		}
	}

	return updatedSkills, nil
}

func (app *UpdateUserAppService) updateCareers(ctx context.Context, userID userdm.UserID, careersReq []updateCareerRequest) ([]*userdm.Career, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	existingCareers, err := app.userRepo.FindCareersByUserID(ctx, userID.String())
	if err != nil {
		return nil, err
	}

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
				userID.String(),
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
				userID.String(),
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
