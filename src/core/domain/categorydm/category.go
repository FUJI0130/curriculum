package categorydm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type Category struct {
	id        CategoryID         `db:"id"`
	name      string             `db:"name"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

func ReconstructCategory(id CategoryID, name string, createdAt time.Time, updatedAt time.Time) (*Category, error) {
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
	return &Category{
		id:        id,
		name:      name,
		createdAt: createdAtByVal,
		updatedAt: updatedAtByVal,
	}, nil
}

func (c *Category) ID() CategoryID {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) CreatedAt() sharedvo.CreatedAt {
	return c.createdAt
}

func (c *Category) UpdatedAt() sharedvo.UpdatedAt {
	return c.updatedAt
}
