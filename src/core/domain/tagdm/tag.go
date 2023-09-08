package tagdm

import (
	"errors"
	"fmt"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
)

type Tag struct {
	id        TagID              `db:"id"`
	name      string             `db:"name"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

// エラーメッセージの追加
var (
	ErrTagNameEmpty = errors.New("tag name cannot be empty")
	ErrTagNotFound  = errors.New("tag not found")
)

func NewTag(name string, created_at time.Time, updatedAt time.Time) (*Tag, error) {

	if name == "" {
		return nil, ErrTagNameEmpty
	}

	tagID, err := NewTagID()
	if err != nil {
		return nil, err
	}
	tagName := name
	tagCreatedAt := sharedvo.NewCreatedAt()
	if err != nil {
		return nil, err
	}

	tagUpdatedAt := sharedvo.NewUpdatedAt()
	if err != nil {
		fmt.Printf("NewTag Time taken for updatedAt.Before(time.Now()): %v\n", sharedvo.LastDuration)
		return nil, err
	}
	return &Tag{
		id:        tagID,
		name:      tagName,
		createdAt: tagCreatedAt,
		updatedAt: tagUpdatedAt,
	}, nil
}

func ReconstructTag(id TagID, name string, created_at time.Time, updatedAt time.Time) (*Tag, error) {
	// タグの名前が空の場合はエラー
	if name == "" {
		return nil, ErrTagNameEmpty
	}

	return &Tag{
		id:        id,
		name:      name,
		createdAt: sharedvo.CreatedAt(created_at),
		updatedAt: sharedvo.UpdatedAt(updatedAt),
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
