package tagdm

import (
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
		return nil, customerrors.ErrTagNameEmpty(nil, "GenWhenCreateTag  name is empty")
	}
	if utf8.RuneCountInString(name) > constants.NameMaxLength {
		return nil, customerrors.ErrTagNameTooLong(nil, "GenWhenCreateTag Tag name is Too long")
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
		return nil, customerrors.ErrTagNameEmpty(nil, "TestNewTag  name is empty")
	}
	if utf8.RuneCountInString(name) > constants.NameMaxLength {
		return nil, customerrors.ErrTagNameTooLong(nil, "TestNewTag Tag name is Too long")
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
