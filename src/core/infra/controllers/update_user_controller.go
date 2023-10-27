package controllers

import (
	"context"
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	UpdateUserService *userapp.UpdateUserAppService
}

func NewUpdateUserController(s *userapp.UpdateUserAppService) *UpdateUserController {
	return &UpdateUserController{UpdateUserService: s}
}

func (ctrl UpdateUserController) Update(c *gin.Context) {
	var req userapp.UpdateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.Error(customerrors.WrapUnprocessableEntityError(err, "JSON binding error"))
		return
	}

	txObj, ok := c.Get("Conn")
	if !ok || txObj == nil {
		c.Error(errors.New("transaction not found"))
		return
	}

	ctxWithTx := context.WithValue(c.Request.Context(), "Conn", txObj)

	if err := ctrl.UpdateUserService.ExecUpdate(ctxWithTx, &req); err != nil {
		if customErr, ok := err.(customerrors.BaseError); ok {
			c.Status(customErr.StatusCode())
		} else {
			c.Status(http.StatusInternalServerError)
		}

		c.Error(customerrors.WrapUnprocessableEntityError(err, "UpdateUserService ExecUpdate error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Updated successfully"})
}

func (ctrl UpdateUserController) Fetch(c *gin.Context) {
	var req userapp.UpdateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.Error(customerrors.WrapUnprocessableEntityError(err, "JSON binding error"))
		return
	}

	txObj, ok := c.Get("Conn")
	if !ok || txObj == nil {
		c.Error(errors.New("transaction not found"))
		return
	}

	ctxWithTx := context.WithValue(c.Request.Context(), "Conn", txObj)

	userID := c.Query("userID") // "userID" はクエリパラメータのキー名

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID is required"})
		return
	}

	user, skills, careers, err := ctrl.UpdateUserService.ExecFetch(ctxWithTx, userID)
	if err != nil {
		if customErr, ok := err.(customerrors.BaseError); ok {
			c.Status(customErr.StatusCode())
		} else {
			c.Status(http.StatusInternalServerError)
		}

		c.Error(customerrors.WrapUnprocessableEntityError(err, "UpdateUserService ExecFetch error"))
		return
	}

	response := map[string]interface{}{
		"UserInfo": map[string]interface{}{
			"Name":     user.Name,
			"Email":    user.Email,
			"Password": user.Password, // Be cautious about sending passwords, you might not want to do this in a real application!
			"Profile":  user.Profile,
		},
		"Skills":  skills,
		"Careers": careers,
	}

	c.JSON(http.StatusOK, response)
}
