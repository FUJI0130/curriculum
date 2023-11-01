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

type UpdateUserRequest struct {
	Command    string `json:"command"`
	UserID     string `json:"userID"`
	UpdateData struct {
		UserInfo UserRequestUpdate      `json:"UserInfo"`
		Skills   []SkillRequestUpdate   `json:"Skills"`
		Careers  []CareersRequestUpdate `json:"Careers"`
	} `json:"updateData"`
}

func (app *UpdateUserAppService) ExecUpdate(ctx context.Context, req *UpdateUserRequest) error {
	log.Printf("req is : %v\n", req)
	log.Printf("req.UserInfo.Email is : %s\n", req.UpdateData.UserInfo.Email)
	userDataOnDB, err := app.userRepo.FindByEmail(ctx, req.UpdateData.UserInfo.Email)
	if err != nil {
		return customerrors.NewNotFoundErrorf("Update_user_app_service  Exec User Not Exist  email is : %s", req.UpdateData.UserInfo.Email)
	}
	log.Printf("log checkpoint1")
	if err = app.updateUsersInfo(ctx, userDataOnDB, req); err != nil {
		return err
	}
	log.Print("log checkpoint2")
	if err = app.updateSkillsInfo(ctx, userDataOnDB, req.UpdateData.Skills); err != nil {
		return err
	}
	log.Print("log checkpoint3")
	if err = app.updateCareersInfo(ctx, userDataOnDB, req.UpdateData.Careers); err != nil {
		return err
	}

	return nil
}

func (app *UpdateUserAppService) updateUsersInfo(ctx context.Context, userDataOnDB *userdm.User, req *UpdateUserRequest) error {
	updatedUser, err := userdm.ReconstructUser(
		userDataOnDB.ID().String(),
		req.UpdateData.UserInfo.Name,
		req.UpdateData.UserInfo.Email,
		req.UpdateData.UserInfo.Password,
		req.UpdateData.UserInfo.Profile,
		userDataOnDB.CreatedAt().DateTime(),
	)
	if err != nil {
		return err
	}

	userDiffs := userDataOnDB.MismatchedFields(updatedUser)
	if len(userDiffs) == 0 {
		return nil
	}

	return app.userRepo.UpdateUser(ctx, updatedUser)
}

func (app *UpdateUserAppService) updateSkillsInfo(ctx context.Context, userDataOnDB *userdm.User, skillsReq []SkillRequestUpdate) error {
	tagNames := make([]string, 0, len(skillsReq))
	seenSkills := make(map[string]bool)
	for _, s := range skillsReq {
		if seenSkills[s.TagName] {
			return customerrors.NewUnprocessableEntityErrorf("Skill with tag name %s is duplicated", s.TagName)
		}
		seenSkills[s.TagName] = true
		tagNames = append(tagNames, s.TagName)
	}

	tags, err := app.tagRepo.FindByNames(ctx, tagNames)
	if err != nil {
		return err
	}

	tagsMap := make(map[string]*tagdm.Tag)
	for _, tag := range tags {
		tagsMap[tag.Name()] = tag
	}

	for _, tagName := range tagNames {
		if _, exists := tagsMap[tagName]; !exists {
			tag, err := tagdm.GenWhenCreateTag(tagName)
			if err != nil {
				return err
			}
			if err = app.tagRepo.Store(ctx, tag); err != nil {
				return err
			}
			tagsMap[tagName] = tag
		}
	}

	//TODO: FindSkillsByUserIDの結果を格納するSkillsのValue Objectが必要
	skills, err := app.userRepo.FindSkillsByUserID(ctx, userDataOnDB.ID().String())
	if err != nil {
		log.Printf("Error fetching skills by userID: %s. Error: %v", userDataOnDB.ID().String(), err)
		return err
	}

	// Log all skills from DB
	log.Printf("Skills from DB:")
	for _, skill := range skills {
		log.Printf("SkillID: %s, TagID: %s, UserID: %s, Evaluation: %d, Years: %d", skill.ID().String(), skill.TagID().String(), skill.UserID().String(), skill.Evaluation().Value(), skill.Year().Value())
	}

	skillsParams := make([]userdm.SkillParam, len(skillsReq))
	for i, s := range skillsReq {
		skillsParams[i] = userdm.SkillParam{
			TagID:      tagsMap[s.TagName].ID(),
			TagName:    s.TagName,
			Evaluation: s.Evaluation,
			Years:      s.Years,
		}
	}

	log.Printf("Skills from Request:")
	for _, s := range skillsParams {
		log.Printf("TagID: %s, TagName: %s, Evaluation: %d, Years: %d", s.TagID.String(), s.TagName, s.Evaluation, s.Years)
	}

	// DBの既存のSkillテーブルのデータと比較して、変更があった場合は変更を行う
	for _, skill := range skills {
		found := false
		for _, s := range skillsParams {
			if skill.TagID().String() == s.TagID.String() {
				found = true

				updateSkill, err := userdm.ReconstructSkill(skill.ID().String(), skill.TagID().String(), skill.UserID().String(), s.Evaluation, s.Years, skill.CreatedAt().DateTime(), time.Now())
				if err != nil {
					return err
				}
				diffs := skill.MismatchedFields(updateSkill)

				if len(diffs) > 0 {
					log.Printf("Skill data is difference")
					newTagID := skill.TagID().String()
					newEvaluation := skill.Evaluation().Value()
					newYears := skill.Year().Value()

					if _, exists := diffs["evaluation"]; exists {
						log.Printf("evaluation is difference")
						newEvaluation = s.Evaluation
					}
					if _, exists := diffs["years"]; exists {
						log.Printf("years is difference")
						newYears = s.Years
					}

					resultSkill, err := userdm.ReconstructSkill(
						skill.ID().String(),
						newTagID,
						skill.UserID().String(),
						newEvaluation,
						newYears,
						skill.CreatedAt().DateTime(),
						skill.UpdatedAt().DateTime(),
					)
					if err != nil {
						return err
					}

					if err = app.userRepo.UpdateSkill(ctx, resultSkill); err != nil {
						return err
					}
				}
				break
			}
		}
		if !found {
			return err
		}
	}
	return err
}

func (app *UpdateUserAppService) updateCareersInfo(ctx context.Context, userDataOnDB *userdm.User, careersReq []CareersRequestUpdate) error {
	careers, err := app.userRepo.FindCareersByUserID(ctx, userDataOnDB.ID().String())
	if err != nil {
		return err
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	careersParams := make([]userdm.CareerParam, len(careersReq))
	for i, c := range careersReq {
		log.Printf("AdFrom is : %s\n", c.AdFrom)
		log.Printf("AdTo is : %s\n", c.AdTo)

		adFromInJST := c.AdFrom.In(jst)
		adToInJST := c.AdTo.In(jst)
		careersParams[i] = userdm.CareerParam{
			Detail: c.Detail,
			AdFrom: adFromInJST,
			AdTo:   adToInJST,
		}
	}

	for _, career := range careers {
		found := false
		for _, c := range careersParams {
			careerAdFrom := toUTCAndTruncate(career.AdFrom())
			careerAdTo := toUTCAndTruncate(career.AdTo())
			cAdFrom := toUTCAndTruncate(c.AdFrom)
			cAdTo := toUTCAndTruncate(c.AdTo)

			log.Printf("DB career adFrom: %s, request career adFrom: %s", careerAdFrom, cAdFrom)
			log.Printf("DB career adTo: %s, request career adTo: %s", careerAdTo, cAdTo)

			if careerAdFrom == cAdFrom {
				log.Printf("AdFrom is same")
			} else {
				log.Printf("AdFrom is not same career.AdFrom is : %s, c.AdFrom is : %s", careerAdFrom, cAdFrom)
			}
			if careerAdTo == cAdTo {
				log.Printf("AdTo is same")
			} else {
				log.Printf("AdTo is not same career.AdTo is : %s, c.AdTo is : %s", careerAdTo, cAdTo)
			}

			if career.Detail() != c.Detail || career.AdFrom() != c.AdFrom || career.AdTo() != c.AdTo {
				found = true

				log.Printf("career checkpoint2")
				updateCareer, err := userdm.ReconstructCareer(
					career.ID().String(),
					c.Detail,
					c.AdFrom,
					c.AdTo,
					career.UserID().String(),
					career.CreatedAt().DateTime(),
					career.UpdatedAt().DateTime(),
				)
				if err != nil {
					return err
				}

				diffs := career.MismatchedFields(updateCareer)

				log.Printf("career checkpoint3")
				if len(diffs) > 0 {

					log.Printf("career checkpoint4")
					newDetail := career.Detail()
					newAdFrom := career.AdFrom()
					newAdTo := career.AdTo()

					if _, exists := diffs["detail"]; exists {
						newDetail = c.Detail
					}
					if _, exists := diffs["adFrom"]; exists {
						newAdFrom = c.AdFrom
					}
					if _, exists := diffs["adTo"]; exists {
						newAdTo = c.AdTo
					}

					finalCareer, err := userdm.ReconstructCareer(
						career.ID().String(),
						newDetail,
						newAdFrom,
						newAdTo,
						career.UserID().String(),
						career.CreatedAt().DateTime(),
						career.UpdatedAt().DateTime(),
					)
					if err != nil {
						return err
					}

					log.Printf("career checkpoint4")
					// データベースに更新を適用
					if err = app.userRepo.UpdateCareer(ctx, finalCareer); err != nil {
						return err
					}
				}
				break
			}
		}
		if !found {
			// return customerrors.NewNotFoundErrorf("Career with UserID %s not found. userEmail is: %s", userDataOnDB.ID().String(), userDataOnDB.Email().String())
			return err
		}
	}
	return err
}
func toUTCAndTruncate(t time.Time) time.Time {
	return t.UTC().Truncate(time.Second)
}

func (app *UpdateUserAppService) ExecFetch(ctx context.Context, userID string) (*userdm.User, []userdm.Skill, []userdm.Career, error) {
	userDataOnDB, err := app.userRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, nil, nil, err
	}

	skills, err := app.userRepo.FindSkillsByUserID(ctx, userID)
	if err != nil {
		return nil, nil, nil, err
	}

	careers, err := app.userRepo.FindCareersByUserID(ctx, userID)
	if err != nil {
		return nil, nil, nil, err
	}

	return userDataOnDB, skills, careers, nil
}
