package categorydm

import (
	"time"
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type CategoryParam struct {
	Name string
}

const CategoryNameMaxLength = 30 // カテゴリ名の最大長を設定します

func GenWhenCreateCategory(categoryID CategoryID, name string, createdAt time.Time) (*Category, error) {
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("Category name is empty")
	}
	if utf8.RuneCountInString(name) > CategoryNameMaxLength {
		return nil, customerrors.NewUnprocessableEntityError("Category name is too long")
	}
	if categoryID == "" {
		return nil, customerrors.NewUnprocessableEntityError("Category ID is empty")
	}

	categoryCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}
	categoryUpdatedAt := sharedvo.NewUpdatedAt() // 現在の時刻で更新日時を生成
	return &Category{
		id:        categoryID,
		name:      name,
		createdAt: categoryCreatedAt,
		updatedAt: categoryUpdatedAt,
	}, nil
}

func TestNewCategory(id string, name string) (*Category, error) {
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("name is empty")
	}
	if utf8.RuneCountInString(name) > CategoryNameMaxLength {
		return nil, customerrors.NewUnprocessableEntityError("Category name is too long")
	}
	categoryID, err := NewCategoryIDFromString(id)
	if err != nil {
		return nil, err
	}
	categoryCreatedAt := sharedvo.NewCreatedAt()
	categoryUpdatedAt := sharedvo.NewUpdatedAt()

	return &Category{
		id:        categoryID,
		name:      name,
		createdAt: categoryCreatedAt,
		updatedAt: categoryUpdatedAt,
	}, nil
}
