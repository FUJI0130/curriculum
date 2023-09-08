package userapp

import (
	"context"
	"errors"
	"fmt"
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
var ErrTagNameAlreadyExists = errors.New("tag name already exists")

// create_user_controller.goのcreateの中で呼び出されてる
func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	existingUser, err := app.userRepo.FindByName(ctx, req.Name)

	if err != nil {
		return err
	}

	if existingUser != nil {
		return ErrUserNameAlreadyExists
	}

	user, err := userdm.NewUser(req.Name, req.Email, req.Password, req.Profile, time.Now(), time.Now())
	if err != nil {
		return err
	}

	log.Printf("before skill Proc")
	skillsSlice := make([]*userdm.Skill, len(req.Skills))
	var tag *tagdm.Tag
	for i, s := range req.Skills {
		// タグ名からタグを検索
		log.Printf("s.TagName is : %s", s.TagName)
		tag, err = app.tagRepo.FindByName(ctx, s.TagName)
		log.Printf("after FindByName")
		if err != nil {
			return err
		}
		log.Printf("after FindByNames error handling")
		// タグが見つからなかった場合、新規タグを作成
		if errors.Is(err, tagdm.ErrTagNotFound) {
			tag, err = app.tagRepo.CreateNewTag(ctx, s.TagName)
			if err != nil {
				return fmt.Errorf("error creating new tag: %v", err)
			}
		}

		skill, err := userdm.NewSkill(tag.ID(), user.ID(), s.Evaluation, s.Years, time.Now(), time.Now())
		if err != nil {
			return fmt.Errorf("error creating new skill: %v", err)
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
