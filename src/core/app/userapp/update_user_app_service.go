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
	Name     string
	Email    string
	Password string
	Skills   []SkillRequest
	Profile  string
	Careers  []CareersRequest
}

func (app *UpdateUserAppService) Exec(ctx context.Context, req *UpdateUserRequest) (err error) {
	userDataOnDB, err := app.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		// emailが一致しなかったらエラーを返す
		return customerrors.NewNotFoundErrorf("Update_user_app_service  Exec User Not Exist  email is : %s", req.Email)
	}

	// Userオブジェクトを再構築
	updatedUser, err := userdm.ReconstructUser(userDataOnDB.ID().String(), userDataOnDB.Name(), req.Email, userDataOnDB.Password().String(), req.Profile, userDataOnDB.CreatedAt().DateTime())

	userDiffs := userDataOnDB.MismatchedFields(updatedUser) // 前提: reqがuserDataOnDBと比較可能で、MismatchedFieldsメソッドが存在する

	if len(userDiffs) > 0 {
		// 変更があった場合のみ新しい値を使用
		newUsername := userDataOnDB.Name()
		newEmail := userDataOnDB.Email().String()
		newPassword := userDataOnDB.Password().String()
		newProfile := userDataOnDB.Profile()

		if _, exists := userDiffs["username"]; exists {
			newUsername = req.Name // reqが該当のフィールドを持っていることを想定
		}
		if _, exists := userDiffs["email"]; exists {
			newEmail = req.Email
		}
		if _, exists := userDiffs["password"]; exists {
			newPassword = req.Password // reqが該当のフィールドを持っていることを想定
		}
		if _, exists := userDiffs["profile"]; exists {
			newProfile = req.Profile // reqが該当のフィールドを持っていることを想定
		}

		// ReconstructUser(id string, name string, email string, password string, profile string, createdAt time.Time) (*User, error)
		// Userオブジェクトを再構築
		updatedUser, err := userdm.ReconstructUser(
			userDataOnDB.ID().String(),
			newUsername,
			newEmail,
			newPassword,
			newProfile,
			userDataOnDB.CreatedAt().DateTime(),
		)
		if err != nil {
			return err
		}

		// データベースに更新を適用
		if err = app.userRepo.UpdateUser(ctx, updatedUser); err != nil {
			return err
		}
	}

	tagNames := make([]string, 0, len(req.Skills))
	seenSkills := make(map[string]bool)
	for _, s := range req.Skills {
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

	skillsParams := make([]userdm.SkillParam, len(req.Skills))
	for i, s := range req.Skills {
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
			return customerrors.NewNotFoundErrorf("Skill with TagID %s not found   userEmail is : %s ", skill.TagID().String(), req.Email)
		}
	}

	careers, err := app.userRepo.FindCareersByUserID(ctx, userDataOnDB.ID().String())

	careersParams := make([]userdm.CareerParam, len(req.Careers))
	for i, c := range req.Careers {
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
			return customerrors.NewNotFoundErrorf("Career with UserID %s not found. userEmail is: %s", userDataOnDB.ID().String(), req.Email)
		}
	}

	return err
}
