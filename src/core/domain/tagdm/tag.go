package tagdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type Tag struct {
	id        TagID
	name      string
	createdAt sharedvo.CreatedAt
	updatedAt sharedvo.UpdatedAt
}

const tagNameMaxLength = 50 // タグ名の最大文字数を定義

func NewTag(name string) (*Tag, error) {
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("tag name is empty")
	}

	if len(name) > tagNameMaxLength {
		return nil, customerrors.NewUnprocessableEntityError("tag name is too long")
	}

	tagID, err := NewTagID() // 新しいタグIDを生成
	if err != nil {
		return nil, customerrors.WrapInternalServerError(err, "failed to generate new tag ID")
	}

	createdAt := sharedvo.NewCreatedAt()
	updatedAt := sharedvo.NewUpdatedAt()

	return &Tag{
		id:        tagID,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}
func ReconstructTag(id TagID, name string, createdAt time.Time, updatedAt time.Time) (*Tag, error) {
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
