//go:generate mockgen -destination=../../mock/mock_mentorrecruitment/mentor_recruitment_repository_mock.go -package=mock_mentorrecruitmentdm . MentorRecruitmentRepository

package mentorrecruitmentdm

import (
	"context"

	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
)

type MentorRecruitmentRepository interface {
	Store(ctx context.Context, mentorRecruitment *MentorRecruitment) error
	FindByID(ctx context.Context, id MentorRecruitmentID) (*MentorRecruitment, error)
	FindByTitle(ctx context.Context, title string) ([]*MentorRecruitment, error)
	FindByCategoryID(ctx context.Context, categoryID categorydm.CategoryID) ([]*MentorRecruitment, error)
	Update(ctx context.Context, mentorRecruitment *MentorRecruitment) error
}
