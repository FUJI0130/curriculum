package userapp

import (
	"context"
	"fmt"
	"time"

	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
	domainErrors "github.com/FUJI0130/curriculum/src/core/domain/customerrors"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/utils"
	"github.com/FUJI0130/curriculum/src/core/validator"
	"github.com/cockroachdb/errors"
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

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	// Validate request keys
	reqMap, err := utils.StructToMap(req)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to convert struct to map. ErrorCode: %d", errorcodes.InternalServerError)
		return errors.Wrapf(err, errMsg)
	}
	if err := validator.ValidateKeysAgainstStruct(reqMap, &CreateUserRequest{}); err != nil {
		errMsg := fmt.Sprintf("Validation failed. ErrorCode: %d", errorcodes.BadRequest)
		return errors.Wrapf(err, errMsg)
	}

	isExist, err := app.existService.Exec(ctx, req.Name)

	//ここのエラーハンドリングcauseに差し替え
	if err != nil {
		errMsg := fmt.Sprintf("Database error. ErrorCode: %d. Failed to check existence of user name: %v", errorcodes.InternalServerError, err)
		return errors.Wrap(err, errMsg)
	}

	if isExist {
		return domainErrors.ErrUserNameAlreadyExists(nil, req.Name, "Create_user_app_service  Exec")
	}

	// 全てのタグ名を一度に取得するためのスライスの作成
	tagNames := make([]string, len(req.Skills))
	for i, s := range req.Skills {
		tagNames[i] = s.TagName
	}

	// 一度のクエリでタグを取得
	tags, err := app.tagRepo.FindByNames(ctx, tagNames)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch tags. ErrorCode: %d", errorcodes.InternalServerError)
		return errors.Wrap(err, errMsg)
	}

	tagsMap := make(map[string]*tagdm.Tag)
	for _, tag := range tags {
		tagsMap[tag.Name()] = tag
	}

	seenSkills := make(map[string]bool)
	skillsParams := make([]userdm.SkillParam, len(req.Skills))

	for i, s := range req.Skills {

		if seenSkills[s.TagName] {
			return domainErrors.ErrDuplicateSkillTag(nil, s.TagName, "Create_user_app_service  Exec")
		}
		seenSkills[s.TagName] = true

		// タグが存在しない場合は新しく作成
		if _, ok := tagsMap[s.TagName]; !ok {
			tag, err := tagdm.GenWhenCreateTag(s.TagName)
			if err != nil {
				errMsg := fmt.Sprintf("Failed to create tag. ErrorCode: %d", errorcodes.InternalServerError)
				return errors.Wrap(err, errMsg)
			}

			if err = app.tagRepo.Store(ctx, tag); err != nil {
				errMsg := fmt.Sprintf("Failed to store tag. ErrorCode: %d", errorcodes.InternalServerError)
				return errors.Wrap(err, errMsg)
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
		errMsg := fmt.Sprintf("Failed to create user. ErrorCode: %d", errorcodes.InternalServerError)
		return errors.Wrap(err, errMsg)
	}

	return app.userRepo.Store(ctx, userdomain)
}
