//go:generate mockgen -destination=../../mock/mock_user/user_repository_mock.go -package=mock_userdm . UserRepository

package userdm

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Store(ctx context.Context, userdomain *UserDomain) error
	StoreWithTransaction(tx *sqlx.Tx, userdomain *UserDomain) error
	FindByName(ctx context.Context, name string) (*User, error)
	FindByNames(ctx context.Context, names []string) (map[string]*User, error)
}
