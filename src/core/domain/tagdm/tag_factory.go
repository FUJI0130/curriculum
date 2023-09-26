package tagdm

import (
	"errors"
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/customerrors"
	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm/constants"
)

type TagParam struct {
	Name string
}

// GenWhenCreateTag creates a new tag with the given parameters.
func GenWhenCreateTag(name string) (*Tag, error) {
	if name == "" {
		return nil, customerrors.ErrTagNameEmpty() // カスタムエラーに差し替え
	}
	if utf8.RuneCountInString(name) > constants.NameMaxLength {
		return nil, customerrors.ErrTagNameTooLong() // カスタムエラーに差し替え
	}
	tagID, err := NewTagID()
	if err != nil {
		return nil, err
	}
	tagCreatedAt := sharedvo.NewCreatedAt()
	tagUpdatedAt := sharedvo.NewUpdatedAt()

	return &Tag{
		id:        tagID,
		name:      name,
		createdAt: tagCreatedAt,
		updatedAt: tagUpdatedAt,
	}, nil
}

// TestNewTag is a function for testing purposes and creates a tag with predefined values.
func TestNewTag(id string, name string) (*Tag, error) {
	if name == "" {
		return nil, customerrors.ErrTagNameEmpty() // カスタムエラーに差し替え
	}
	if utf8.RuneCountInString(name) > constants.NameMaxLength {
		return nil, customerrors.ErrTagNameTooLong() // カスタムエラーに差し替え
	}
	tagID, err := NewTagIDFromString(id)
	if err != nil {
		return nil, err
	}
	tagCreatedAt := sharedvo.NewCreatedAt()
	tagUpdatedAt := sharedvo.NewUpdatedAt()

	return &Tag{
		id:        tagID,
		name:      name,
		createdAt: tagCreatedAt,
		updatedAt: tagUpdatedAt,
	}, nil
}

var ErrTagTooLong = errors.New("tag name is too long")
