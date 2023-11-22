// mentor_recruitment_query_repository.go
package mentorrecruitmentdm

import (
	"context"

	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
)

type MentorRecruitmentQueryRepository interface {
	FindByID(ctx context.Context, id MentorRecruitmentID) (*MentorRecruitment, error)
	FindByCategoryID(ctx context.Context, categoryID categorydm.CategoryID) ([]*MentorRecruitment, error)
}
