package datamodel

import "time"

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Profile   string    `db:"profile"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Skill struct {
	ID         string    `db:"id"`
	TagID      string    `db:"tag_id"`
	UserID     string    `db:"user_id"`
	Evaluation uint8     `db:"evaluation"`
	Years      uint8     `db:"years"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
type Career struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Detail    string    `db:"detail"`
	AdFrom    time.Time `db:"ad_from"`
	AdTo      time.Time `db:"ad_to"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
