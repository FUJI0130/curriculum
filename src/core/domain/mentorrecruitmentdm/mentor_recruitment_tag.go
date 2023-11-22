package mentorrecruitmentdm

import (
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
)

type MentorRecruitmentTag struct {
	MentorRecruitmentID MentorRecruitmentID
	TagID               tagdm.TagID
	CreatedAt           time.Time
}

func NewMentorRecruitmentTag(mentorRecruitmentID MentorRecruitmentID, tagID tagdm.TagID, createdAt time.Time) *MentorRecruitmentTag {
	return &MentorRecruitmentTag{
		MentorRecruitmentID: mentorRecruitmentID,
		TagID:               tagID,
		CreatedAt:           createdAt,
	}
}
