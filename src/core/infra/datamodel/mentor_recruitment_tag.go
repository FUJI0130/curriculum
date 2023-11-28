package datamodel

import "time"

type MentorRecruitmentTag struct {
	MentorRecruitmentTagID string    `db:"mentor_recruitment_id"`
	TagID                  string    `db:"tag_id"`
	CreatedAt              time.Time `db:"created_at"`
}
