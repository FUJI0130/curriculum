// - `Store() error`というインターフェースでユーザ作成をするので、作っておく
// - また`FindByName() (*User, error)`は名前の重複チェックで使うので入れておく

package userdm

import "context"

type UserRepository interface {
	Store(ctx context.Context, user *User) error
	FindByName(ctx context.Context, name string) (*User, error)
}
