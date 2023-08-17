package userapp

import (
	"context"
	"errors"
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
	Name         string
	Introduction string
	Password     string
	Skills       []string
	Experiences  []string
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

	id := userdm.NewUserID()
	user, err := userdm.NewUser(id, req.Name, req.Password, req.Introduction, req.Skills, req.Experiences)
	if err != nil {
		return err
	}

	return app.userRepo.Store(user)
}
