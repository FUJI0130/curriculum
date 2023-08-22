package userapp

import (
	"context"
	"errors"
	"time"

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

// Controllerからユースケースへの引数は必ずプリミティブで渡すこと！
// 理由はドメインをなるべくユースケース以下の円の外側(オニオンアーキテクチャの円を参照)の部分で
// 触らせないようにすることで、ドメイン知識が漏れ出すのを防ぐため
type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
	Skills   []string
	Profile  string
	Careers  []string
}

// ユーザドメイン作成
// ユーザ名重複チェック
// ユーザ作成
func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	existingUser, err := app.userRepo.FindByName(req.Name)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user name already exists")
	}

	// createdAt, err := sharedvo.NewCreatedAt(time.Now())
	// updatedAt, err := sharedvo.NewUpdatedAt(time.Now())
	// func NewUser(name string, email string, password string, profile string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	user, err := userdm.NewUser(req.Name, req.Email, req.Password, req.Profile, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return app.userRepo.Store(user)
}
