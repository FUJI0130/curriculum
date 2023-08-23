package userapp

import (
	"context"
	"errors"
	"log"
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

// var ErrUserNotFound = errors.New("user not found")
var ErrUserNameAlreadyExists = errors.New("user name already exists")

// ユーザドメイン作成
// ユーザ名重複チェック
// ユーザ作成
func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	existingUser, err := app.userRepo.FindByName(ctx, req.Name)
	if err != nil {
		if errors.Is(err, userdm.ErrUserNotFound) {
			// ユーザーが存在しないので、新しいユーザーの作成を続行する
		} else {
			log.Println("Error after FindByName:", err)
			return err
		}
	}

	if existingUser != nil {
		// existingUser が nil ではない場合、ユーザー名が既に存在すると判断
		log.Println("Existing user details:", existingUser)
		return ErrUserNameAlreadyExists
	}

	user, err := userdm.NewUser(req.Name, req.Email, req.Password, req.Profile, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return app.userRepo.Store(ctx, user)
}
