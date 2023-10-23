//go:generate mockgen -destination=../../mock/mock_tag/tag_repository_mock.go -package=mock_tagdm . TagRepository

package tagdm

import (
	"context"

	"github.com/FUJI0130/curriculum/src/core/domain"
)

type TagRepository interface {
	Store(ctx context.Context, tag *Tag) error
	StoreWithTransaction(transaction domain.Transaction, tag *Tag) error
	FindByName(ctx context.Context, name string) (*Tag, error)
	FindByNames(ctx context.Context, names []string) ([]*Tag, error)
	FindByID(ctx context.Context, id string) (*Tag, error)
}
