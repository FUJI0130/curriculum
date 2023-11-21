//go:generate mockgen -destination=../../mock/mock_user/user_repository_mock.go -package=mock_userdm . UserRepository

package userdm

import (
	"context"
)

type UserRepository interface {
	Store(ctx context.Context, userEntity *User) error
	FindByUserName(ctx context.Context, name string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUserID(ctx context.Context, userID string) (*User, error)
	Update(ctx context.Context, userdomain *User) error
}
