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
	userID    UserID             `db:"user_id"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

func NewCareer(detail string, adFromSet time.Time, adToSet time.Time, userID UserID) (*Career, error) {

	if detail == "" { // 例えば、detailが空の場合のエラーハンドリング
		return nil, customerrors.NewUnprocessableEntityError("[NewCareer]  detail is empty")
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
		userID:    userID,
		createdAt: careerCreatedAt,
		updatedAt: careerUpdatedAt,
	}, nil
}

// 以下のゲッターはそのままです。
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

func (c *Career) UserID() UserID {
	return c.userID
}

func (c *Career) CreatedAt() sharedvo.CreatedAt {
	return c.createdAt
}

func (c *Career) UpdatedAt() sharedvo.UpdatedAt {
	return c.updatedAt
}
