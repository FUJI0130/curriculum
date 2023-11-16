//go:generate mockgen -destination=../../mock/mock_user/user_repository_mock.go -package=mock_userdm . UserRepository

package userdm

import (
	"context"
)

type UserRepository interface {
	Store(ctx context.Context, userdomain *UserDomain) error
	FindByUserName(ctx context.Context, name string) (*UserDomain, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUserID(ctx context.Context, userID string) (*UserDomain, error)
	Update(ctx context.Context, userdomain *UserDomain) error
}
