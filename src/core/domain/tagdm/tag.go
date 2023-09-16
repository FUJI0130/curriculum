package tagdm

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
)

type Tag struct {
	id        TagID              `db:"id"`
	name      string             `db:"name"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

var (
	ErrTagNameEmpty = errors.New("tag name cannot be empty")
	ErrTagNotFound  = errors.New("tag not found")
)

const nameMaxLength = 15

func NewTag(name string) (*Tag, error) {

	if name == "" {
		return nil, ErrTagNameEmpty
	}
	if utf8.RuneCountInString(name) > nameMaxLength {
		return nil, fmt.Errorf("タグの名前が最大許容長の%dを超えています", nameMaxLength)
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

func ReconstructTag(id TagID, name string, createdAt time.Time, updatedAt time.Time) (*Tag, error) {
	if name == "" {
		return nil, ErrTagNameEmpty
	}
	createdAtByVal, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}
	updatedAtByVal, err := sharedvo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		return nil, err
	}
	return &Tag{
		id:        id,
		name:      name,
		createdAt: createdAtByVal,
		updatedAt: updatedAtByVal,
	}, nil
}

func (t *Tag) ID() TagID {
	return t.id
}
func (t *Tag) Name() string {
	return t.name
}

func (t *Tag) CreatedAt() sharedvo.CreatedAt {
	return t.createdAt
}

func (t *Tag) UpdatedAt() sharedvo.UpdatedAt {
	return t.updatedAt
}
