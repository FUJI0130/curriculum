// src/core/app/mentorapp/get_mentor_list_app_service.go

package mentorapp

import "github.com/FUJI0130/curriculum/src/core/domain/mentordm"

type GetMentorListAppService struct {
	mentorRepo mentordm.MentorRepository
}

func NewGetMentorListAppService(mentorRepo mentordm.MentorRepository) *GetMentorListAppService {
	return &GetMentorListAppService{mentorRepo}
}

func (service *GetMentorListAppService) Execute() ([]*mentordm.Mentor, error) {
	mentors, err := service.mentorRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return mentors, nil
}
