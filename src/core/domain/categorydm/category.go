package categorydm

import (
	"time"
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type Category struct {
	id        CategoryID         `db:"id"`
	name      string             `db:"name"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

func NewCategory(name string) (*Category, error) {
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("Category name is empty")
	}
	if utf8.RuneCountInString(name) > CategoryNameMaxLength {
		return nil, customerrors.NewUnprocessableEntityError("Category name is too long")
	}

	categoryID, err := NewCategoryID() // カテゴリIDを生成する関数（実装が必要）
	if err != nil {
		return nil, err
	}

	createdAt := sharedvo.NewCreatedAt() // 現在の時刻で作成日時を生成
	updatedAt := sharedvo.NewUpdatedAt() // 現在の時刻で更新日時を生成

	return &Category{
		id:        categoryID,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}
func ReconstructCategory(id CategoryID, name string, createdAt time.Time, updatedAt time.Time) (*Category, error) {
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("name is empty")
	}
	createdAtByVal, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityError(err, "createdAt is invalid")
	}
	updatedAtByVal, err := sharedvo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityError(err, "updatedAt is invalid")
	}
	return &Category{
		id:        id,
		name:      name,
		createdAt: createdAtByVal,
		updatedAt: updatedAtByVal,
	}, nil
}

func (c *Category) ID() CategoryID {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) CreatedAt() sharedvo.CreatedAt {
	return c.createdAt
}

func (c *Category) UpdatedAt() sharedvo.UpdatedAt {
	return c.updatedAt
}
