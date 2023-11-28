// mentor_recruitment_command_repository.go
package mentorrecruitmentdm

import (
	"context"
)

type MentorRecruitmentCommandRepository interface {
	Store(ctx context.Context, mentorRecruitment *MentorRecruitment) error
}
