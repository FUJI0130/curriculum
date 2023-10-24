//go:generate mockgen -destination=../../mock/mock_user/user_repository_mock.go -package=mock_userdm . UserRepository

package userdm

import (
	"context"
)

type UserRepository interface {
	Store(ctx context.Context, userdomain *UserDomain) error
	FindByName(ctx context.Context, name string) (*User, error)
	FindByNames(ctx context.Context, names []string) (map[string]*User, error)
}
