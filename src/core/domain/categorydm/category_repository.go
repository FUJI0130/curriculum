//go:generate mockgen -destination=../../mock/mock_categorydm/category_repository_mock.go -package=mock_categorydm . CategoryRepository

package categorydm

import (
	"context"
)

type CategoryRepository interface {
	Store(ctx context.Context, category *Category) error
	FindByName(ctx context.Context, name string) (*Category, error)
	FindByID(ctx context.Context, id string) (*Category, error)
}
