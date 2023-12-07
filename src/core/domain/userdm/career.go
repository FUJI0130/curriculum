package userdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type Career struct {
	id        CareerID           `db:"id"`
	detail    string             `db:"detail"`
	adFrom    time.Time          `db:"ad_from"`
	adTo      time.Time          `db:"ad_to"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

func NewCareer(detail string, adFromSet time.Time, adToSet time.Time) (*Career, error) {

	if detail == "" {
		return nil, customerrors.NewUnprocessableEntityError("Career Detail is empty")
	}

	careerID, err := NewCareerID()
	if err != nil {
		return nil, err
	}

	careerCreatedAt := sharedvo.NewCreatedAt()
	if err != nil {
		return nil, err
	}

	careerUpdatedAt := sharedvo.NewUpdatedAt()
	if err != nil {
		return nil, err
	}

	return &Career{
		id:        careerID,
		detail:    detail,
		adFrom:    adFromSet,
		adTo:      adToSet,
		createdAt: careerCreatedAt,
		updatedAt: careerUpdatedAt,
	}, nil
}

func (c *Career) ID() CareerID {
	return c.id
}

func (c *Career) Detail() string {
	return c.detail
}

func (c *Career) AdFrom() time.Time {
	return c.adFrom
}

func (c *Career) AdTo() time.Time {
	return c.adTo
}

func (c *Career) CreatedAt() sharedvo.CreatedAt {
	return c.createdAt
}

func (c *Career) UpdatedAt() sharedvo.UpdatedAt {
	return c.updatedAt
}

func ReconstructCareer(id string, detail string, adFrom time.Time, adTo time.Time, createdAt time.Time, updatedAt time.Time) (*Career, error) {
	careerID, err := NewCareerIDFromString(id)
	if err != nil {
		return nil, err
	}

	careerCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}

	careerUpdatedAt, err := sharedvo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		return nil, err
	}

	return &Career{
		id:        careerID,
		detail:    detail,
		adFrom:    adFrom,
		adTo:      adTo,
		createdAt: careerCreatedAt,
		updatedAt: careerUpdatedAt,
	}, nil
}

func (c *Career) Equal(other *Career) bool {
	if other == nil {
		return false
	}

	return c.id.Equal(other.id) &&
		c.detail == other.detail &&
		c.adFrom == other.adFrom &&
		c.adTo == other.adTo &&
		c.createdAt.Equal(other.createdAt) &&
		c.updatedAt.Equal(other.updatedAt)
}
