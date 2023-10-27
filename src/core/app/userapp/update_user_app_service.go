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

type UpdateUserRequest struct {
	UserInfo UserRequest
	Skills   []SkillRequest
	Careers  []CareersRequest
}

func (app *UpdateUserAppService) ExecUpdate(ctx context.Context, req *UpdateUserRequest) error {
	userDataOnDB, err := app.userRepo.FindByEmail(ctx, req.UserInfo.Email)
	if err != nil {
		return customerrors.NewNotFoundErrorf("Update_user_app_service  Exec User Not Exist  email is : %s", req.UserInfo.Email)
	}

	if err = app.updateUserInformation(ctx, userDataOnDB, req); err != nil {
		return err
	}

	if err = app.updateSkills(ctx, userDataOnDB, req.Skills); err != nil {
		return err
	}

	if err = app.updateCareers(ctx, userDataOnDB, req.Careers); err != nil {
		return err
	}

	return nil
}

func (app *UpdateUserAppService) updateUserInformation(ctx context.Context, userDataOnDB *userdm.User, req *UpdateUserRequest) error {
	updatedUser, err := userdm.ReconstructUser(
		userDataOnDB.ID().String(),
		req.UserInfo.Name,
		req.UserInfo.Email,
		req.UserInfo.Password,
		req.UserInfo.Profile,
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

func (app *UpdateUserAppService) updateSkills(ctx context.Context, userDataOnDB *userdm.User, skillsReq []SkillRequest) error {
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

	skillsParams := make([]userdm.SkillParam, len(skillsReq))
	for i, s := range skillsReq {
		skillsParams[i] = userdm.SkillParam{
			TagID:      tagsMap[s.TagName].ID(),
			TagName:    s.TagName,
			Evaluation: s.Evaluation,
			Years:      s.Years,
		}
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

				// 変更があった場合、Skillオブジェクトを再構築して更新
				if len(diffs) > 0 {
					// 既存のスキルの値をデフォルトとして使用
					newTagID := skill.TagID().String()
					newEvaluation := skill.Evaluation().Value() // 前提としてEvaluation()がuint8を返すか、Value()メソッドを持っていることを想定
					newYears := skill.Year().Value()            // 前提としてYears()がuint8を返すか、Value()メソッドを持っていることを想定

					// 変更があった場合のみ新しい値を使用
					if _, exists := diffs["Evaluation"]; exists {
						newEvaluation = s.Evaluation
					}
					if _, exists := diffs["Years"]; exists {
						newYears = s.Years
					}

					// Skillオブジェクトを再構築
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

					// データベースに更新を適用
					if err = app.userRepo.UpdateSkill(ctx, resultSkill); err != nil {
						return err
					}
				}
				break
			}
		}
		if !found {
			return customerrors.NewNotFoundErrorf("Skill with TagID %s not found   userEmail is : %s ", skill.TagID().String(), userDataOnDB.Email().String())
		}
	}
	return err
}

func (app *UpdateUserAppService) updateCareers(ctx context.Context, userDataOnDB *userdm.User, careersReq []CareersRequest) error {
	careers, err := app.userRepo.FindCareersByUserID(ctx, userDataOnDB.ID().String())
	if err != nil {
		return err
	}

	careersParams := make([]userdm.CareerParam, len(careersReq))
	for i, c := range careersReq {
		careersParams[i] = userdm.CareerParam{
			Detail: c.Detail,
			AdFrom: c.AdFrom,
			AdTo:   c.AdTo,
		}
	}

	//DBの既存のCareerテーブルのデータと比較して、変更があった場合は変更を行う
	for _, career := range careers {
		found := false
		for _, c := range careersParams {
			if career.Detail() == c.Detail && career.AdFrom() == c.AdFrom && career.AdTo() == c.AdTo {
				found = true

				// Careerオブジェクトを再構築
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

				if len(diffs) > 0 {

					// 変更があった場合のみ新しい値を使用
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

					// Careerオブジェクトを再構築して変更を反映
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

					// データベースに更新を適用
					if err = app.userRepo.UpdateCareer(ctx, finalCareer); err != nil {
						return err
					}
				}
				break
			}
		}
		if !found {
			return customerrors.NewNotFoundErrorf("Career with UserID %s not found. userEmail is: %s", userDataOnDB.ID().String(), userDataOnDB.Email().String())
		}
	}
	return err
}

func (app *UpdateUserAppService) ExecFetch(ctx context.Context, userID string) (*userdm.User, []userdm.Skill, []userdm.Career, error) {
	userDataOnDB, err := app.userRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, nil, nil, customerrors.NewNotFoundErrorf("Update_user_app_service  ExecFetch User Not Exist  userID is : %s", userID)
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
