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
}

func NewCreateUserAppService(userRepo userdm.UserRepository) *CreateUserAppService {
	return &CreateUserAppService{
		userRepo: userRepo,
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
	TagID      string
	UserID     string
	Evaluation uint8
	Years      uint8
}

type CareersRequest struct {
	Detail string
	AdFrom time.Time
	AdTo   time.Time
	UserID string
}

var ErrUserNameAlreadyExists = errors.New("user name already exists")

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	existingUser, err := app.userRepo.FindByName(ctx, req.Name)
	if err != nil {
		if errors.Is(err, userdm.ErrUserNotFound) {
		} else {
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

	// TODO:　後程タグIDの部分は修正する　一時的にこの形で渡すこととする 23/8/30
	tagId, err := tagdm.NewTagID()
	if err != nil {
		return err
	}
	skillsSlice := make([]*userdm.Skills, len(req.Skills))
	for i, s := range req.Skills {

		// func NewSkills(tagId tagdm.TagID, userId UserID, evaluation uint8, years uint8) (*Skills, error) {
		skill, err := userdm.NewSkills(tagId, user.ID(), s.Evaluation, s.Years)
		if err != nil {
			return err
		}
		skillsSlice[i] = skill
	}

	// careersSlice の作成
	careersSlice := make([]*userdm.Careers, len(req.Careers))
	for i, c := range req.Careers {
		career, err := userdm.NewCareers(c.Detail, c.AdFrom, c.AdTo, user.ID(), user.CreatedAt().DateTime(), user.UpdatedAt().DateTime())
		if err != nil {
			return err
		}
		careersSlice[i] = career
	}
	//結局ここにもcareersSliceが必要な気がしてきた

	return app.userRepo.Store(ctx, user, skillsSlice, careersSlice)
}
