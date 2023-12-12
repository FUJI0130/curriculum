package router

import (
	"github.com/FUJI0130/curriculum/src/core/app/mentorapp"
	"github.com/FUJI0130/curriculum/src/core/infra/controllers"
	"github.com/gin-gonic/gin"
)

func InitMentorRoutes(r *gin.Engine, createMentorRecruitmentService *mentorapp.CreateMentorRecruitmentAppService) {
	InitCreateMentorRecruitmentController(r, createMentorRecruitmentService)
}

func InitCreateMentorRecruitmentController(r *gin.Engine, s *mentorapp.CreateMentorRecruitmentAppService) {
	ctrl := controllers.NewCreateMentorRecruitmentController(s)
	r.POST("/mentor_recruitments", ctrl.Create)
}
