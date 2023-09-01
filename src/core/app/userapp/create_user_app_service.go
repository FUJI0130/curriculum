package userapp

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
)

type CreateUserAppService struct {
	userRepo userdm.UserRepository
	tagRepo  tagdm.TagRepository
}

func NewCreateUserAppService(userRepo userdm.UserRepository, tagRepo tagdm.TagRepository) *CreateUserAppService {
	return &CreateUserAppService{
		userRepo: userRepo,
		tagRepo:  tagRepo,
	}
}

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
	Skills   []SkillRequest
	Profile  string
	Careers  []CareersRequest
}

type SkillRequest struct {
	TagName    string
	Evaluation uint8
	Years      uint8
}

type CareersRequest struct {
	Detail string
	AdFrom time.Time
	AdTo   time.Time
}

var ErrUserNameAlreadyExists = errors.New("user name already exists")

// create_user_controller.goのcreateの中で呼び出されてる
func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	existingUser, err := app.userRepo.FindByName(ctx, req.Name)
	if err != nil {
		if !errors.Is(err, userdm.ErrUserNotFound) {
			log.Println("Error after FindByName:", err)
			return err
		}
	}

	if existingUser != nil {
		log.Println("Existing user details:", existingUser)
		return ErrUserNameAlreadyExists
	}

	user, err := userdm.NewUser(req.Name, req.Email, req.Password, req.Profile, time.Now(), time.Now())
	if err != nil {
		return err
	}

	skillsSlice := make([]*userdm.Skill, len(req.Skills))

	var tag *tagdm.Tag
	for i, s := range req.Skills {
		// 名前からタグを検索
		tag, err = app.tagRepo.FindByName(ctx, s.TagName)

		// タグが見つからない場合
		if errors.Is(err, tagdm.ErrTagNotFound) {
			// 新規タグを作成 (このメソッドはあなたのリポジトリ設計による)
			tag, err = app.tagRepo.CreateNewTag(ctx, s.TagName) // 仮にTagNameがタグの名前を示すものとして
			if err != nil {
				return err // 新規タグの作成中にエラーが発生した場合
			}
		} else if err != nil {
			return err // その他のエラー
		}

		// tag.ID() を使用してスキルを作成
		skill, err := userdm.NewSkill(tag.ID(), user.ID(), s.Evaluation, s.Years)
		if err != nil {
			return err
		}
		skillsSlice[i] = skill
	}

	// careersSlice の作成
	careersSlice := make([]*userdm.Career, len(req.Careers))
	for i, c := range req.Careers {
		career, err := userdm.NewCareer(c.Detail, c.AdFrom, c.AdTo, user.ID(), user.CreatedAt().DateTime(), user.UpdatedAt().DateTime())
		if err != nil {
			return err
		}
		careersSlice[i] = career
	}

	return app.userRepo.Store(ctx, user, skillsSlice, careersSlice)
}
