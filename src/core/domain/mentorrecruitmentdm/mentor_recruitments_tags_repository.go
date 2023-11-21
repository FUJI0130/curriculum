package mentorrecruitmentdm

import (
	"context"
)

type MentorRecruitmentsTagsRepository interface {
	Store(ctx context.Context, mentorRecruitmentTag *MentorRecruitmentTag) error
	FindByMentorRecruitmentID(ctx context.Context, mentorRecruitmentID MentorRecruitmentID) ([]*MentorRecruitmentTag, error)
}
