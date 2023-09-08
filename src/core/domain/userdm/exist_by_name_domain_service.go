package userdm

import (
	"context"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
)

type existByNameDomainService struct {
	userRepo UserRepository
}

// TODO: コンストラクタ作る

func (ds *existByNameDomainService) Exec(ctx context.Context, name string) (bool, error) {
	// 元々書いてある処理を書く
	existingUser, err := userapp.userRepo.FindByName(ctx, name)

	if err != nil {
		return err
	}

	if existingUser != nil {
		return userapp.ErrUserNameAlreadyExists
	}
}
