package tagdm

import (
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type TagParam struct {
	Name string
}

const NameMaxLength = 15

func GenWhenCreateTag(name string) (*Tag, error) {
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("Tag name is empty")
	}
	if utf8.RuneCountInString(name) > NameMaxLength {
		return nil, customerrors.NewUnprocessableEntityError("Tag name is Too long")
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

func TestNewTag(id string, name string) (*Tag, error) {
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("name is empty")
	}
	if utf8.RuneCountInString(name) > NameMaxLength {
		return nil, customerrors.NewUnprocessableEntityError("Tag name is Too long")
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
