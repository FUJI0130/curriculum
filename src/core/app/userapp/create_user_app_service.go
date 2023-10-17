package userapp

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/jmoiron/sqlx"
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

	reqMap, err := StructToMap(req)
	if err != nil {
		return err
	}
	if err := ValidateKeysAgainstStruct(reqMap, &CreateUserRequest{}); err != nil {
		return err
	}

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

func (app *CreateUserAppService) ExecWithTransaction(ctx context.Context, tx *sqlx.Tx, req *CreateUserRequest) error {

	// Validate request keys
	reqMap, err := StructToMap(req)
	if err != nil {
		// return customerrors.WrapInternalServerError(err, "Failed to convert struct to map. ")
		return err
	}
	if err := ValidateKeysAgainstStruct(reqMap, &CreateUserRequest{}); err != nil {
		// return customerrors.WrapUnprocessableEntityError(err, "Validation failed. ")
		return err
	}

	isExist, err := app.existService.Exec(ctx, req.Name)

	if err != nil {
		// return customerrors.WrapInternalServerErrorf(err, "Database error.  Failed to check existence of user name: %s", req.Name)
		return err
	}

	if isExist {
		// return customerrors.NewUnprocessableEntityErrorf("Create_user_app_service  Exec UserName isExist  name is : %s", req.Name)
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
		// return customerrors.WrapInternalServerError(err, "Failed to fetch tags. ")
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
			// return customerrors.NewUnprocessableEntityErrorf("Create_user_app_service  Exec tagname is : %s", s.TagName)
			return err
		}
		seenSkills[s.TagName] = true

		// タグが存在しない場合は新しく作成
		if _, ok := tagsMap[s.TagName]; !ok {
			tag, err := tagdm.GenWhenCreateTag(s.TagName)
			if err != nil {
				// return customerrors.WrapInternalServerError(err, "Failed to create tag. ")
				return err
			}

			if err = app.tagRepo.StoreWithTransaction(tx, tag); err != nil {
				// return customerrors.WrapInternalServerError(err, "Failed to store tag. ")
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
		// return customerrors.WrapInternalServerError(err, "Failed to create user. ")
		return err
	}

	return app.userRepo.StoreWithTransaction(tx, userdomain)
}

func StructToMap(req interface{}) (map[string]any, error) {
	result := make(map[string]interface{})
	val := reflect.ValueOf(req)
	if val.Kind() != reflect.Ptr {
		return nil, customerrors.NewInternalServerError("StructToMap Expected a pointer but got " + val.Kind().String())
	}
	val = val.Elem() // Now it's safe to call Elem
	typ := val.Type()

	// Iterate over struct fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Ignore unexported fields
		if fieldType.PkgPath != "" {
			continue
		}

		key := fieldType.Name

		// Recursively process nested structures
		if field.Kind() == reflect.Struct {
			nestedMap, err := StructToMap(field.Addr().Interface())
			if err != nil {
				return nil, err
			}
			result[key] = nestedMap
		} else if field.Kind() == reflect.Slice {
			var sliceData []interface{}
			for si := 0; si < field.Len(); si++ {
				sliceElem := field.Index(si)
				if sliceElem.Kind() == reflect.Struct {
					nestedMap, err := StructToMap(sliceElem.Addr().Interface())
					if err != nil {
						return nil, err
					}
					sliceData = append(sliceData, nestedMap)
				} else {
					sliceData = append(sliceData, sliceElem.Interface())
				}
			}
			result[key] = sliceData
		} else {
			// Plain field, just set the value
			result[key] = field.Interface()
		}
	}
	return result, nil
}

func ValidateKeysAgainstStruct(rawReq map[string]interface{}, referenceStruct any) error {
	expectedKeys := make(map[string]bool)

	val := reflect.ValueOf(referenceStruct).Elem()
	for i := 0; i < val.NumField(); i++ {
		expectedKeys[val.Type().Field(i).Name] = true
	}

	for key, value := range rawReq {
		// Check if key is expected
		if _, exists := expectedKeys[key]; !exists {
			return customerrors.NewUnprocessableEntityError(fmt.Sprintf("Unexpected key '%s' in the request", key))
		}

		// Recursively check nested structures
		if nestedMap, ok := value.(map[string]interface{}); ok {
			field, _ := val.Type().FieldByName(key)
			if field.Type.Kind() == reflect.Struct {
				if err := ValidateKeysAgainstStruct(nestedMap, reflect.New(field.Type).Interface()); err != nil {
					return customerrors.NewUnprocessableEntityError(fmt.Sprintf("key '%s': %v", key, err))
				}
			}
		} else if nestedSlice, ok := value.([]interface{}); ok {
			field, _ := val.Type().FieldByName(key)
			if field.Type.Elem().Kind() == reflect.Struct {
				for i, nestedElement := range nestedSlice {
					if nestedMap, ok := nestedElement.(map[string]interface{}); ok {
						if err := ValidateKeysAgainstStruct(nestedMap, reflect.New(field.Type.Elem()).Interface()); err != nil {
							return customerrors.NewUnprocessableEntityError(fmt.Sprintf("key '%s' index %d: %v", key, i, err))
						}
					}
				}
			}
		}
	}
	return nil
}
