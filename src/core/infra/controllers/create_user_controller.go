package controllers

import (
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	createUserService *userapp.CreateUserAppService
}

func NewCreateUserController(s *userapp.CreateUserAppService) *CreateUserController {
	return &CreateUserController{createUserService: s}
}

func (ctrl *CreateUserController) Create(c *gin.Context) {
	var req userapp.CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.Error(customerrors.WrapUnprocessableEntityError(err, "JSON binding error"))
		return
	}

	if err := ctrl.createUserService.Exec(c.Request.Context(), &req); err != nil {
		if customErr, ok := err.(customerrors.BaseError); ok {
			c.Status(customErr.StatusCode())
		} else {
			c.Status(http.StatusInternalServerError)
		}

		c.Error(customerrors.WrapUnprocessableEntityError(err, "createUserService Exec error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
