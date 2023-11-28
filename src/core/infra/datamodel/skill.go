package datamodel

import "time"

type Skill struct {
	ID         string    `db:"id"`
	TagID      string    `db:"tag_id"`
	UserID     string    `db:"user_id"`
	Evaluation uint8     `db:"evaluation"`
	Years      uint8     `db:"years"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
