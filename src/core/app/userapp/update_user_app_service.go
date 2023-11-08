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
	Command    string `json:"command"`
	UserID     string `json:"userID"`
	UpdateData struct {
		Users   UpdateUserRequest     `json:"Users"`
		Skills  []UpdateSkillRequest  `json:"Skills"`
		Careers []UpdateCareerRequest `json:"Careers"`
	} `json:"updateData"`
}

func (app *UpdateUserAppService) ExecUpdate(ctx context.Context, req *UpdateUserRequestData) error {
	log.Printf("req is : %v\n", req)
	log.Printf("req.Users.Email is : %s\n", req.UpdateData.Users.Email)
	userDataOnDB, err := app.userRepo.FindByUserID(ctx, req.UpdateData.Users.ID)
	if err != nil {
		return customerrors.NewNotFoundErrorf("Update_user_app_service  Exec User Not Exist  email is : %s", req.UpdateData.Users.Email)
	}
	log.Printf("log checkpoint1")
	if err = app.updateUsers(ctx, userDataOnDB, &req.UpdateData.Users); err != nil {
		return err
	}
	log.Print("log checkpoint2")
	if err = app.updateSKills(ctx, userDataOnDB, req.UpdateData.Skills); err != nil {
		return err
	}
	log.Print("log checkpoint3")
	if err = app.updateCareers(ctx, userDataOnDB, req.UpdateData.Careers); err != nil {
		return err
	}

	return nil
}

func (app *UpdateUserAppService) updateUsers(ctx context.Context, userDataOnDB *userdm.User, req *UpdateUserRequest) error {
	updatedUser, err := userdm.GenUserWhenUpdate(
		userDataOnDB.ID().String(),
		req.Name,
		req.Email,
		req.Password,
		req.Profile,
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

func (app *UpdateUserAppService) updateSKills(ctx context.Context, userDataOnDB *userdm.User, skillsReq []UpdateSkillRequest) error {
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
			//N+1
			if err = app.tagRepo.Store(ctx, tag); err != nil {
				return err
			}
			tagsMap[tagName] = tag
		}
	}

	skills, err := app.userRepo.FindSkillsByUserID(ctx, userDataOnDB.ID().String())
	if err != nil {
		return err
	}

	existingSkillsMap := make(map[string]*userdm.Skill)
	for _, skill := range skills {
		log.Printf("Processing skill with ID: %s and TagID: %s", skill.ID().String(), skill.TagID().String()) // ← ログ追加: 現在のスキルのIDとTagIDを出力
		skillCopy := skill
		existingSkillsMap[skill.TagID().String()] = &skillCopy
	}

	// ログ追加: existingSkillsMapの全体的なサイズと中身を出力
	log.Printf("Size of existingSkillsMap: %d", len(existingSkillsMap))
	for tagID, skill := range existingSkillsMap {
		log.Printf("Key (TagID): %s, Skill ID: %s", tagID, skill.ID().String())
	}

	for _, s := range skillsReq {
		tagID := tagsMap[s.TagName].ID().String()
		log.Printf("Processing skill with tagID: %s", tagID) // ← ログ追加

		if skill, exists := existingSkillsMap[tagID]; exists {
			log.Println("Updating existing skill") // ← ログ追加
			log.Printf("GenSkillWhenUpdate Params:\nSkillID: %s\nTagID: %s\nUserID: %s\nEvaluation: %d\nYears: %d\nCreatedAt: %v\nUpdatedAt: %v\n",
				skill.ID().String(), tagID, skill.UserID().String(), s.Evaluation, s.Years, skill.CreatedAt().DateTime(), time.Now())

			storeSkill, err := userdm.GenSkillWhenUpdate(skill.ID().String(), tagID, skill.UserID().String(), s.Evaluation, s.Years, skill.CreatedAt().DateTime(), time.Now())
			log.Println("after GenSkillWhenUpdate")
			if err != nil {
				return err
			}
			if err = app.userRepo.StoreSkill(ctx, storeSkill); err != nil {
				return err
			}
			log.Println("after StoreSkill")
		} else {
			log.Println("Creating new skill") // ← ログ追加
			// Handle scenario where a new skill is present in the request but not in the DB
			newSkill, err := userdm.GenSkillWhenCreate(
				tagID,
				userDataOnDB.ID().String(),
				s.Evaluation,
				s.Years,
				time.Now(), // Current time as createdAt
				time.Now(), // Current time as updatedAt
			)
			if err != nil {
				return err
			}

			// Assuming you have a method to store a new skill in the userRepo
			if err = app.userRepo.StoreSkill(ctx, newSkill); err != nil {
				return err
			}
		}
	}

	return nil
}

func (app *UpdateUserAppService) updateCareers(ctx context.Context, userDataOnDB *userdm.User, careersReq []UpdateCareerRequest) error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// ユーザーIDに基づいて既存のキャリア情報を取得
	existingCareers, err := app.userRepo.FindCareersByUserID(ctx, userDataOnDB.ID().String())
	if err != nil {
		return err
	}

	existingCareersMap := make(map[string]*userdm.Career)
	for _, career := range existingCareers {
		careerCopy := career
		existingCareersMap[career.ID().String()] = &careerCopy
	}

	for _, c := range careersReq {
		adFromInJST := c.AdFrom.In(jst)
		adToInJST := c.AdTo.In(jst)

		if career, exists := existingCareersMap[c.ID]; exists {
			// 既存のキャリア情報がリクエストに存在する場合、内容を置き換え
			updatedCareer, err := userdm.GenCareerWhenUpdate(
				career.ID().String(),
				c.Detail,
				adFromInJST,
				adToInJST,
				userDataOnDB.ID().String(),
				career.CreatedAt().DateTime(),
				time.Now(), // Current time as updatedAt
			)
			if err != nil {
				return err
			}
			if err = app.userRepo.StoreCareer(ctx, updatedCareer); err != nil {
				return err
			}
		} else {
			// リクエストに新しいキャリア情報があれば、それを追加
			newCareer, err := userdm.GenCareerWhenCreate(
				c.Detail,
				adFromInJST,
				adToInJST,
				userDataOnDB.ID().String(),
				time.Now(), // Current time as createdAt
				time.Now(), // Current time as updatedAt
			)
			if err != nil {
				return err
			}
			if err = app.userRepo.StoreCareer(ctx, newCareer); err != nil {
				return err
			}
		}
	}

	return nil
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
