package userdm

import (
	"fmt"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
)

type Careers struct {
	id        CareerID           `db:"id"`
	detail    string             `db:"detail"`
	adFrom    time.Time          `db:"ad_from"`
	adTo      time.Time          `db:"ad_to"`
	userId    UserID             `db:"user_id"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

func NewCareers(detail string, adFrom_Set time.Time, adTo_Set time.Time, userId UserID, created_at time.Time, updatedAt time.Time) (*Careers, error) {

	career_id, err := NewCareerID()
	if err != nil {
		return nil, err
	}

	careerCreatedAt, err := sharedvo.NewCreatedAt(created_at)
	if err != nil {
		return nil, err
	}

	careerUpdatedAt, err := sharedvo.NewUpdatedAt(updatedAt)
	if err != nil {
		fmt.Printf("Time taken for updatedAt.Before(time.Now()): %v\n", sharedvo.LastDuration)
		return nil, err
	}

	return &Careers{
		id:        career_id,
		detail:    detail,
		adFrom:    adFrom_Set,
		adTo:      adTo_Set,
		userId:    userId,
		createdAt: *careerCreatedAt,
		updatedAt: *careerUpdatedAt,
	}, nil
}

// 以下のゲッターはそのままです。
func (c *Careers) ID() CareerID {
	return c.id
}

func (c *Careers) Detail() string {
	return c.detail
}

func (c *Careers) AdFrom() time.Time {
	return c.adFrom
}

func (c *Careers) AdTo() time.Time {
	return c.adTo
}

func (c *Careers) UserID() UserID {
	return c.userId
}

func (c *Careers) CreatedAt() sharedvo.CreatedAt {
	return c.createdAt
}

func (c *Careers) UpdatedAt() sharedvo.UpdatedAt {
	return c.updatedAt
}
