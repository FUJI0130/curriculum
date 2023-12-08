package controllers

import (
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/app/mentorapp"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
)

type CreateMentorRecruitmentController struct {
	createMentorRecruitmentService *mentorapp.CreateMentorRecruitmentAppService
}

func NewCreateMentorRecruitmentController(s *mentorapp.CreateMentorRecruitmentAppService) *CreateMentorRecruitmentController {
	return &CreateMentorRecruitmentController{createMentorRecruitmentService: s}
}

func (ctrl *CreateMentorRecruitmentController) Create(c *gin.Context) {

	var req mentorapp.CreateMentorRecruitmentRequest
	if err := c.BindJSON(&req); err != nil {
		c.Error(customerrors.WrapUnprocessableEntityError(err, "JSON binding error"))
		return
	}

	if err := ctrl.createMentorRecruitmentService.Exec(c.Request.Context(), &req); err != nil {
		if customErr, ok := err.(customerrors.BaseError); ok {
			c.Status(customErr.StatusCode())
		} else {
			c.Status(http.StatusInternalServerError)
		}

		c.Error(customerrors.WrapUnprocessableEntityError(err, "createMentorRecruitmentService Exec error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mentor recruitment created successfully"})
}
