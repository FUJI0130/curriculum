// src/infra/controllers/mentor_controller.go

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/app/mentorapp"
)

type MentorController struct {
	getMentorListAppService *mentorapp.GetMentorListAppService
}

func NewMentorController(getMentorListAppService *mentorapp.GetMentorListAppService) *MentorController {
	return &MentorController{getMentorListAppService}
}

func (controller *MentorController) GetMentorList(w http.ResponseWriter, r *http.Request) {
	mentors, err := controller.getMentorListAppService.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(mentors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
