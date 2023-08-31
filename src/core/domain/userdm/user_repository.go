// - `Store() error`というインターフェースでユーザ作成をするので、作っておく
// - また`FindByName() (*User, error)`は名前の重複チェックで使うので入れておく
//go:generate mockgen -destination=../../mock/mockUser/user_repository_mock.go -package=mockUser . UserRepository

package userdm

import "context"

type UserRepository interface {
	Store(ctx context.Context, user *User, skills []*Skills, careers []*Careers) error
	FindByName(ctx context.Context, name string) (*User, error)
}

// type Careers struct {
// 	id        CareerID           `db:"id"`
// 	detail    string             `db:"detail"`
// 	adFrom    time.Time          `db:"ad_from"`
// 	adTo      time.Time          `db:"ad_to"`
// 	userId    UserID             `db:"user_id"`
// 	createdAt sharedvo.CreatedAt `db:"created_at"`
// 	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
// }

// type CareersRequest struct {
// 	Id     string
// 	From   uint16
// 	To     uint16
// 	Detail string
// }
