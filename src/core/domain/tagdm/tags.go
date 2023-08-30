package tagdm

import (
	"fmt"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
)

type Tags struct {
	id        TagID              `db:"id"`
	name      string             `db:"name"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

func NewTags(name string, created_at time.Time, updatedAt time.Time) (*Tags, error) {
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
	return &Tags{
		id:        tagId,
		name:      tagName,
		createdAt: *tagCreatedAt,
		updatedAt: *tagUpdatedAt,
	}, nil
}
