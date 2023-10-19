package controllers

import (
	"log"
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const (
	errCodeConflict            = 409
	errCodeInternalServerError = 500
	errCodeNotFound            = 404
	errCodeUnprocessableEntity = 400
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
		c.Error(customerrors.WrapUnprocessableEntityError(err, "create_user_controller [Create] : JSON binding error"))
		return
	}

	if err := ctrl.createUserService.Exec(c, &req); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
func (ctrl *CreateUserController) CreateWithTransaction(c *gin.Context) {
	var req userapp.CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.Error(customerrors.WrapUnprocessableEntityError(err, "create_user_controller [CreateWithTransaction] : JSON binding error"))
		return
	}

	ctx := c.Request.Context()
	tx, _ := c.MustGet("tx").(*sqlx.Tx)

	if err := ctrl.createUserService.ExecWithTransaction(ctx, tx, &req); err != nil {
		if customErr, ok := err.(customerrors.BaseError); ok {
			log.Println("Error is of type customerrors.BaseError")
			c.Error(customErr)
			c.Status(customErr.StatusCode())
			// c.JSON(customErr.StatusCode(), gin.H{"message": customErr.Error()})
		} else {
			log.Println("Error is NOT of type customerrors.BaseError")
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
