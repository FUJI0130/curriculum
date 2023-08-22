// - `Store() error`というインターフェースでユーザ作成をするので、作っておく
// - また`FindByName() (*User, error)`は名前の重複チェックで使うので入れておく

package userdm

type UserRepository interface {
	Store(user *User) error
	FindByName(name string) (*User, error)
}
