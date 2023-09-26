package userapp

import (
	"context"
	"errors"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
)

type CreateUserAppService struct {
	userRepo     userdm.UserRepository
	tagRepo      tagdm.TagRepository
	existService userdm.ExistByNameDomainService
}

func NewCreateUserAppService(userRepo userdm.UserRepository, tagRepo tagdm.TagRepository, existService userdm.ExistByNameDomainService) *CreateUserAppService {
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

var test = ""

var ErrUserNameAlreadyExists = errors.New("user name already exists")
var ErrTagNameAlreadyExists = errors.New("tag name already exists")

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	isExist, err := app.existService.Exec(ctx, req.Name)
	if err != nil {
		return err
	}

	if isExist {
		return ErrUserNameAlreadyExists
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
	tagsMap := make(map[string]*tagdm.Tag)
	for _, tag := range tags {
		tagsMap[tag.Name()] = tag
	}
	seenSkills := make(map[string]bool)
	skillsParams := make([]userdm.SkillParam, len(req.Skills))

	for i, s := range req.Skills {

		if seenSkills[s.TagName] {
			return errors.New("同じスキルタグを複数回持つことはできません")
		}
		seenSkills[s.TagName] = true

		// タグが存在しない場合は新しく作成
		if _, ok := tagsMap[s.TagName]; !ok {
			tag, err := tagdm.GenWhenCreateTag(s.TagName)
			if err != nil {
				return err
			}

			if err = app.tagRepo.Store(ctx, tag); err != nil {
				return err
			}
			tagsMap[s.TagName] = tag
		}

		skillsParams[i] = userdm.SkillParam{
			TagID:      tagsMap[s.TagName].ID(),
			TagName:    s.TagName,
			Evaluation: s.Evaluation,
			Years:      s.Years,
		}

	}

	careersParams := make([]userdm.CareerParam, len(req.Careers))
	for i, c := range req.Careers {
		careersParams[i] = userdm.CareerParam{
			Detail: c.Detail,
			AdFrom: c.AdFrom,
			AdTo:   c.AdTo,
		}
	}

	userdomain, err := userdm.GenWhenCreate(req.Name, req.Email, req.Password, req.Profile, skillsParams, careersParams)
	if err != nil {
		return err
	}

	return app.userRepo.Store(ctx, userdomain)
}
