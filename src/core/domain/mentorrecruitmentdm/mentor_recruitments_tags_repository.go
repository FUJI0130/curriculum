package mentorrecruitmentdm

import (
	"context"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
)

type MentorRecruitmentsTagsRepository interface {
	Store(ctx context.Context, mentorRecruitmentID MentorRecruitmentID, tagID tagdm.TagID) error
	FindByMentorRecruitmentID(ctx context.Context, mentorRecruitmentID MentorRecruitmentID) ([]tagdm.TagID, error)
}
