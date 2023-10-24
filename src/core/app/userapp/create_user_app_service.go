package userapp

import (
	"context"
	"log"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
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

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {

	isExist, err := app.existService.Exec(ctx, req.Name)

	if err != nil {
		return err
	}

	if isExist {
		return customerrors.NewUnprocessableEntityErrorf("Create_user_app_service  Exec UserName isExist  name is : %s", req.Name)
	}

	tagNames := make([]string, len(req.Skills))
	for i, s := range req.Skills {
		tagNames[i] = s.TagName
	}

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

			return customerrors.NewUnprocessableEntityErrorf("Create_user_app_service  Exec tagname is : %s", s.TagName)
		}
		seenSkills[s.TagName] = true

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

func (app *CreateUserAppService) ExecWithTransaction(ctx context.Context, req *CreateUserRequest) (err error) {

	log.Printf("ExecWithTransaction: %s", req.Name)
	isExist, err := app.existService.Exec(ctx, req.Name)
	if err != nil {
		return err
	}

	log.Printf("before isExist: %v", isExist)
	if isExist {
		return customerrors.NewUnprocessableEntityErrorf("Create_user_app_service  Exec UserName isExist  name is : %s", req.Name)
	}
	log.Printf("after isExist: %v", isExist)

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
			if err = app.tagRepo.StoreWithTransaction(ctx, tag); err != nil {
				return err
			}
			tagsMap[tagName] = tag
		}
	}

	skillsParams := make([]userdm.SkillParam, len(req.Skills))
	for i, s := range req.Skills {
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

	log.Printf("before StoreWithTransaction: %s", userdomain.User.Name())
	err = app.userRepo.StoreWithTransaction(ctx, userdomain)

	return err
}
