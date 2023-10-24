package controllers

import (
	"context"
	"net/http"

	"github.com/cockroachdb/errors"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
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
		c.Error(customerrors.WrapUnprocessableEntityError(err, "JSON binding error"))
		return
	}

	txObj, ok := c.Get("transaction")
	if !ok || txObj == nil {
		c.Error(errors.New("transaction not found"))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	ctxWithTx := context.WithValue(c.Request.Context(), "transaction", txObj)

	if err := ctrl.createUserService.Exec(ctxWithTx, &req); err != nil {
		c.Error(customerrors.WrapUnprocessableEntityError(err, "createUserService Exec error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
