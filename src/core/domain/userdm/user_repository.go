// - `Store() error`というインターフェースでユーザ作成をするので、作っておく
// - また`FindByName() (*User, error)`は名前の重複チェックで使うので入れておく
//go:generate mockgen -destination=../../mock/mockUser/user_repository_mock.go -package=mockUser . UserRepository

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
