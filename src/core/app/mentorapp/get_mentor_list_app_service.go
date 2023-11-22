package mentorapp

import (
	"context"

	"github.com/FUJI0130/curriculum/src/core/domain/mentorrecruitmentdm"
)

type GetMentorListAppService struct {
	mentorRecruitmentRepo mentorrecruitmentdm.MentorRecruitmentQueryRepository
}

func NewGetMentorListAppService(mentorRecruitmentRepo mentorrecruitmentdm.MentorRecruitmentQueryRepository) *GetMentorListAppService {
	return &GetMentorListAppService{
		mentorRecruitmentRepo: mentorRecruitmentRepo,
	}
}

func (s *GetMentorListAppService) Execute(ctx context.Context) ([]*mentorrecruitmentdm.MentorRecruitment, error) {
	mentorRecruitments, err := s.mentorRecruitmentRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return mentorRecruitments, nil
}
