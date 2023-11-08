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
	UpdateData struct {
		Users   UpdateUserRequest     `json:"Users"`
		Skills  []UpdateSkillRequest  `json:"Skills"`
		Careers []UpdateCareerRequest `json:"Careers"`
	} `json:"updateData"`
}

func (app *UpdateUserAppService) ExecUpdate(ctx context.Context, req *UpdateUserRequestData) error {
	log.Printf("req is : %v\n", req)

	// ユーザー情報をデータベースから取得
	userDataOnDB, err := app.userRepo.FindByUserID(ctx, req.UpdateData.Users.ID)
	if err != nil {
		return customerrors.NewNotFoundErrorf("Update_user_app_service  Exec User Not Exist  UserID is : %s", req.UpdateData.Users.ID)
	}
	log.Printf("log checkpoint1")

	// ユーザー情報の更新を実行し、更新されたUserエンティティを取得
	updatedUser, err := app.updateUsers(ctx, userDataOnDB, &req.UpdateData.Users)
	if err != nil {
		return err
	}
	log.Print("log checkpoint2")

	// スキル情報の更新を実行し、更新されたSkillエンティティのスライスを取得
	updatedSkills, err := app.updateSkills(ctx, updatedUser, req.UpdateData.Skills)
	if err != nil {
		return err
	}
	log.Print("log checkpoint3")

	// キャリア情報の更新を実行し、更新されたCareerエンティティのスライスを取得
	updatedCareers, err := app.updateCareers(ctx, updatedUser.ID(), req.UpdateData.Careers)
	if err != nil {
		return err
	}

	// UserDomainを構築
	userDomain, err := userdm.GenWhenUpdate(updatedUser, updatedSkills, updatedCareers)
	if err != nil {
		return err
	}

	// リポジトリを使用してUserDomainをデータベースに保存
	if err := app.userRepo.Update(ctx, userDomain); err != nil {
		return err
	}

	return nil
}

func (app *UpdateUserAppService) updateUsers(ctx context.Context, userDataOnDB *userdm.User, req *UpdateUserRequest) (*userdm.User, error) {
	// ユーザー情報の更新オブジェクトを生成
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

	// ユーザー情報に変更がなければ更新をスキップ
	userDiffs := userDataOnDB.MismatchedFields(updatedUser)
	if len(userDiffs) == 0 {
		return userDataOnDB, nil
	}

	// 更新されたユーザーオブジェクトを返す
	return updatedUser, nil
}

func (app *UpdateUserAppService) updateSkills(ctx context.Context, userDataOnDB *userdm.User, skillsReq []UpdateSkillRequest) ([]*userdm.Skill, error) {
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
		skillCopy := skill // Copy the skill
		existingSkillsMap[skill.TagID().String()] = &skillCopy
	}

	for _, s := range skillsReq {
		tagID := tagsMap[s.TagName].ID().String()
		if skill, exists := existingSkillsMap[tagID]; exists {
			// Update the skill
			updatedSkill, err := userdm.GenSkillWhenUpdate(skill.ID().String(), tagID, skill.UserID().String(), s.Evaluation, s.Years, skill.CreatedAt().DateTime(), time.Now())
			if err != nil {
				return nil, err
			}
			updatedSkills = append(updatedSkills, updatedSkill)
		} else {
			// Create a new skill
			newSkill, err := userdm.GenSkillWhenCreate(tagID, userDataOnDB.ID().String(), s.Evaluation, s.Years, time.Now(), time.Now())
			if err != nil {
				return nil, err
			}
			updatedSkills = append(updatedSkills, newSkill)
		}
	}

	return updatedSkills, nil
}

func (app *UpdateUserAppService) updateCareers(ctx context.Context, userID userdm.UserID, careersReq []UpdateCareerRequest) ([]*userdm.Career, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	// 既存のキャリア情報を取得
	existingCareers, err := app.userRepo.FindCareersByUserID(ctx, userID.String())
	if err != nil {
		return nil, err
	}

	// 既存のキャリア情報をIDをキーとしてマップに格納
	existingCareersMap := make(map[string]*userdm.Career)
	for _, career := range existingCareers {
		careerCopy := career
		existingCareersMap[career.ID().String()] = &careerCopy
	}

	var careersToUpdate []*userdm.Career

	for _, c := range careersReq {
		adFromInJST := c.AdFrom.In(jst)
		adToInJST := c.AdTo.In(jst)
		var career *userdm.Career

		// 既存のキャリア情報があるかどうか確認
		if existingCareer, ok := existingCareersMap[c.ID]; ok {
			// 既存のキャリア情報を更新する
			career, err = userdm.GenCareerWhenUpdate(
				existingCareer.ID().String(),
				c.Detail,
				adFromInJST,
				adToInJST,
				userID.String(),
				existingCareer.CreatedAt().DateTime(),
				time.Now(), // UpdatedAtを現在時刻に設定
			)
			if err != nil {
				return nil, err
			}
		} else {
			// 新しいキャリア情報を作成する
			career, err = userdm.GenCareerWhenCreate(
				c.Detail,
				adFromInJST,
				adToInJST,
				userID.String(),
				time.Now(), // CreatedAtは現在時刻
				time.Now(), // UpdatedAtは現在時刻
			)
			if err != nil {
				return nil, err
			}
		}
		careersToUpdate = append(careersToUpdate, career)
	}

	// 更新が必要なキャリアのリストを返す
	return careersToUpdate, nil
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
