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

var ErrTagNotFound = errors.New("tag not found")

func NewTag(name string, created_at time.Time, updatedAt time.Time) (*Tag, error) {
	tagId, err := NewTagID()
	if err != nil {
		return nil, err
	}
	tagName := name
	tagCreatedAt, err := sharedvo.NewCreatedAt(created_at)
	if err != nil {
		return nil, err
	}

	tagUpdatedAt, err := sharedvo.NewUpdatedAt(updatedAt)
	if err != nil {
		fmt.Printf("Time taken for updatedAt.Before(time.Now()): %v\n", sharedvo.LastDuration)
		return nil, err
	}
	return &Tag{
		id:        tagId,
		name:      tagName,
		createdAt: *tagCreatedAt,
		updatedAt: *tagUpdatedAt,
	}, nil
}

// エラーメッセージの追加
var (
	ErrTagNameEmpty = errors.New("tag name cannot be empty")
)

func ReconstructTag(id TagID, name string, created_at time.Time, updatedAt time.Time) (*Tag, error) {
	// タグの名前が空の場合はエラー
	if name == "" {
		return nil, ErrTagNameEmpty
	}

	// 他のエラーチェックもここで行うことができます。

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
