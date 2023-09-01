// - `Store() error`というインターフェースでユーザ作成をするので、作っておく
// - また`FindByName() (*User, error)`は名前の重複チェックで使うので入れておく
//go:generate mockgen -destination=../../mock/mockUser/tag_repository_mock.go -package=mockUser . TagRepository

package tagdm

import "context"

type TagRepository interface {
	Store(ctx context.Context, tag *Tag) error
	FindByName(ctx context.Context, name string) (*Tag, error)
	FindByID(ctx context.Context, id string) (*Tag, error)
	CreateNewTag(ctx context.Context, name string) (*Tag, error)
}
