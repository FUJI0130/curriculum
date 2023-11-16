package categorydm

import (
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type CategoryParam struct {
	Name string
}

const CategoryNameMaxLength = 30 // カテゴリ名の最大長を設定します

func GenWhenCreateCategory(name string) (*Category, error) {
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("Category name is empty")
	}
	if utf8.RuneCountInString(name) > CategoryNameMaxLength {
		return nil, customerrors.NewUnprocessableEntityError("Category name is too long")
	}
	categoryID, err := NewCategoryID()
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
