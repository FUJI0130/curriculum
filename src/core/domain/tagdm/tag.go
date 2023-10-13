package tagdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type Tag struct {
	id        TagID              `db:"id"`
	name      string             `db:"name"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

func ReconstructTag(id TagID, name string, createdAt time.Time, updatedAt time.Time) (*Tag, error) {
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("[ReconstructTag] name is empty")
	}
	createdAtByVal, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityError(err, "[ReconstructTag] createdAt is invalid")
	}
	updatedAtByVal, err := sharedvo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityError(err, "[ReconstructTag] updatedAt is invalid")
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
