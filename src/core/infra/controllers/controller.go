package controllers

import (
	"github.com/FUJI0130/curriculum/src/core/app/mentorapp"
	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/gin-gonic/gin"
)

func InitControllers(r *gin.Engine, createUserService *userapp.CreateUserAppService, updateUserService *userapp.UpdateUserAppService, createMentorRecruitmentService *mentorapp.CreateMentorRecruitmentAppService) {
	InitHealthController(r)
	InitCreateUserController(r, createUserService)
	InitUpdateUserController(r, updateUserService)
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

func InitCreateMentorRecruitmentController(r *gin.Engine, s *mentorapp.CreateMentorRecruitmentAppService) {
	ctrl := NewCreateMentorRecruitmentController(s)
	r.POST("/mentor_recruitments", ctrl.Create)
}
