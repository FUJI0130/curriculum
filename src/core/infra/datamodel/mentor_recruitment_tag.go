package datamodel

import "time"

type MentorRecruitmentTag struct {
	ID                     string    `db:"id"`
	MentorRecruitmentTagID string    `db:"mentor_recruitment_id"`
	TagID                  string    `db:"tag_id"`
	CreatedAt              time.Time `db:"created_at"`
}
