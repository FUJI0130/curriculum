//go:generate mockgen -destination=../../mock/mockTag/tag_repository_mock.go -package=mockTag . TagRepository

package tagdm

import (
	"context"
)

type TagRepository interface {
	Store(ctx context.Context, tag *Tag) error
	FindByName(ctx context.Context, name string) (*Tag, error)
	FindByNames(ctx context.Context, names []string) ([]*Tag, error)
	FindByID(ctx context.Context, id string) (*Tag, error)
}
