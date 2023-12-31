package userapp

import (
	"context"

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

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) (err error) {

	isExist, err := app.existService.Exec(ctx, req.Name)
	if err != nil {
		return err
	}

	if isExist {
		return customerrors.NewUnprocessableEntityErrorf("Create_user_app_service  Exec UserName isExist  name is : %s", req.Name)
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
			//N+1
			if err = app.tagRepo.Store(ctx, tag); err != nil {
				return err
			}
			tagsMap[tagName] = tag
		}
	}

	skills := make([]userdm.Skill, len(req.Skills))
	for i, s := range req.Skills {
		tag, exists := tagsMap[s.TagName]
		if !exists {
			return customerrors.NewUnprocessableEntityErrorf("Tag not found for skill: %s", s.TagName)
		}

		skill, err := userdm.NewSkill(tag.ID(), s.Evaluation, s.Years)
		if err != nil {
			return err
		}
		skills[i] = *skill
	}

	careers := make([]userdm.Career, len(req.Careers))
	for i, c := range req.Careers {
		career, err := userdm.NewCareer(c.Detail, c.AdFrom, c.AdTo)
		if err != nil {
			return err
		}
		careers[i] = *career
	}

	userdomain, err := userdm.GenWhenCreate(req.Name, req.Email, req.Password, req.Profile, skills, careers)
	if err != nil {
		return err
	}

	err = app.userRepo.Store(ctx, userdomain)

	return err
}
