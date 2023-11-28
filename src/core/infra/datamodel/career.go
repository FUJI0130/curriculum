package datamodel

import "time"

type Career struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Detail    string    `db:"detail"`
	AdFrom    time.Time `db:"ad_from"`
	AdTo      time.Time `db:"ad_to"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
