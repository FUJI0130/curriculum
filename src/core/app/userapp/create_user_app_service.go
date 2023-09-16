package userapp

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm/interfaces"
)

type CreateUserAppService struct {
	userRepo     userdm.UserRepository
	tagRepo      tagdm.TagRepository
	existService interfaces.ExistByNameDomainService
}

func NewCreateUserAppService(userRepo userdm.UserRepository, tagRepo tagdm.TagRepository, existService interfaces.ExistByNameDomainService) *CreateUserAppService {
	return &CreateUserAppService{
		userRepo:     userRepo,
		tagRepo:      tagRepo,
		existService: existService,
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

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	isExist, err := app.existService.IsExist(ctx, req.Name)
	if err != nil {
		return err
	}

	if isExist {
		return ErrUserNameAlreadyExists
	}

	userFactory := userdm.NewUserFactory()
	user, err := userFactory.GenWhenCreate(req.Name, req.Email, req.Password, req.Profile)

	if err != nil {
		return err
	}
	// 全てのタグ名を一度に取得するためのスライスの作成
	tagNames := make([]string, len(req.Skills))
	for i, s := range req.Skills {
		tagNames[i] = s.TagName
	}

	// 一度のクエリでタグを取得
	tags, err := app.tagRepo.FindByNames(ctx, tagNames)
	if err != nil {
		return err
	}

	seenSkills := make(map[string]bool)
	skillsSlice := make([]*userdm.Skill, len(req.Skills))
	// var tag *tagdm.Tag
	for i, s := range req.Skills {

		if seenSkills[s.TagName] {
			return errors.New("同じスキルタグを複数回持つことはできません")
		}
		seenSkills[s.TagName] = true

		// tag, err = app.tagRepo.FindByName(ctx, s.TagName)
		// if errors.Is(err, tagdm.ErrTagNotFound) {
		tag, ok := tags[s.TagName]
		// タグが存在しない場合は新しく作成
		if !ok {
			tag, err = tagdm.NewTag(s.TagName)
			if err != nil {
				return err
			}

			if err = app.tagRepo.Store(ctx, tag); err != nil {
				return err
			}
		}
		// } else if err != nil {
		// return err
		// }

		skill, err := userdm.NewSkill(tag.ID(), user.ID(), s.Evaluation, s.Years, time.Now(), time.Now())
		if err != nil {
			return fmt.Errorf("error creating new skill: %v", err)
		}

		skillsSlice[i] = skill
	}

	careersSlice := make([]*userdm.Career, len(req.Careers))
	for i, c := range req.Careers {
		career, err := userdm.NewCareer(c.Detail, c.AdFrom, c.AdTo, user.ID())
		if err != nil {
			return err
		}
		careersSlice[i] = career
	}

	return app.userRepo.Store(ctx, user, skillsSlice, careersSlice)
}
