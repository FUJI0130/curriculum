//go:generate mockgen -destination=../../mock/mock_tag/tag_repository_mock.go -package=mock_tagdm . TagRepository

package tagdm

import (
	"context"
)

type TagRepository interface {
	Store(ctx context.Context, tag *Tag) error
	StoreWithTransaction(ctx context.Context, tag *Tag) error
	FindByName(ctx context.Context, name string) (*Tag, error)
	FindByNames(ctx context.Context, names []string) ([]*Tag, error)
	FindByID(ctx context.Context, id string) (*Tag, error)
}
