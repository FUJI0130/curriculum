package controllers

import (
	"log"

	"github.com/FUJI0130/curriculum/src/core/app/mentorapp"
	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/gin-gonic/gin"
)

func InitControllers(r *gin.Engine, createUserService *userapp.CreateUserAppService, updateUserService *userapp.UpdateUserAppService, createMentorRecruitmentService *mentorapp.CreateMentorRecruitmentAppService) {
	InitHealthController(r)
	InitCreateUserController(r, createUserService)
	InitUpdateUserController(r, updateUserService)
	log.Println("InitCreateMentorRecruitmentController")
	InitCreateMentorRecruitmentController(r, createMentorRecruitmentService)
}

func InitCreateUserController(r *gin.Engine, s *userapp.CreateUserAppService) {
	ctrl := NewCreateUserController(s)
	r.POST("/users", ctrl.Create)
}

func InitUpdateUserController(r *gin.Engine, UpdateuserService *userapp.UpdateUserAppService) {
	ctrl := NewUpdateUserController(UpdateuserService)
	r.PATCH("/users", ctrl.Update)
}

func InitMentorController(r *gin.Engine, getMentorListAppService *mentorapp.GetMentorListAppService) {
	// ctrl := NewMentorController(getMentorListAppService)
	// r.GET("/mentors", ctrl.Fetch)
}
func InitCreateMentorRecruitmentController(r *gin.Engine, s *mentorapp.CreateMentorRecruitmentAppService) {
	log.Println("InitCreateMentorRecruitmentController1")
	ctrl := NewCreateMentorRecruitmentController(s)
	log.Println("InitCreateMentorRecruitmentController2")
	r.POST("/mentor_recruitments", ctrl.Create)
}
