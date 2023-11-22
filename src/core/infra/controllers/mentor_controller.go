// src/infra/controllers/mentor_controller.go

package controllers

import (
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/app/mentorapp"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
)

type MentorController struct {
	getMentorListAppService *mentorapp.GetMentorListAppService
}

func NewMentorController(getMentorListAppService *mentorapp.GetMentorListAppService) *MentorController {
	return &MentorController{getMentorListAppService}
}

func (controller *MentorController) GetMentorList(c *gin.Context) {
	mentors, err := controller.getMentorListAppService.Execute(c.Request.Context())
	if err != nil {
		// エラーに応じて適切なステータスコードを設定
		if customErr, ok := err.(customerrors.BaseError); ok {
			c.Status(customErr.StatusCode())
		} else {
			c.Status(http.StatusInternalServerError)
		}

		// エラー情報をレスポンスに含める
		c.Error(customerrors.WrapInternalServerError(err, "Error executing getMentorListAppService"))
		return
	}

	// 成功した場合、メンターリストをJSON形式で返す
	c.JSON(http.StatusOK, mentors)
}
