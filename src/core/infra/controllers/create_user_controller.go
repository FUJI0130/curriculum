package controllers

import (
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
		c.JSON(customerrors.ErrCodeUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// コンテキストの取得
	ctx := c.Request.Context()

	// トランザクションの取得
	tx, _ := c.MustGet("tx").(*sqlx.Tx)

	// ExecWithTransactionにコンテキストとトランザクションを渡す
	if err := ctrl.createUserService.ExecWithTransaction(ctx, tx, &req); err != nil {
		switch err.(type) {
		case *customerrors.ConflictErrorType:
			c.JSON(customerrors.ErrCodeConflict, gin.H{"error": err.Error()})
		case *customerrors.InternalServerErrorType:
			c.JSON(customerrors.ErrCodeInternalServerError, gin.H{"error": err.Error()})
		case *customerrors.NotFoundErrorType:
			c.JSON(customerrors.ErrCodeNotFound, gin.H{"error": err.Error()})
		case *customerrors.UnprocessableEntityErrorType:
			c.JSON(customerrors.ErrCodeUnprocessableEntity, gin.H{"error": err.Error()})
		default:
			c.JSON(customerrors.ErrCodeInternalServerError, gin.H{"error": err.Error()}) // 予期せぬエラーの場合、500を返す
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
