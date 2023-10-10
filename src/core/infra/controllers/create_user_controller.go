// FWからのリクエストをユースケースに送り、レスポンスを返す
package controllers

import (
	"net/http"

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

// ここでcurlコマンドの内容をバインドしている
func (ctrl *CreateUserController) Create(c *gin.Context) {

	var req userapp.CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(errCodeUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.createUserService.Exec(c, &req); err != nil {
		switch err.(type) {
		case *customerrors.ConflictErrorType:
			c.JSON(errCodeConflict, gin.H{"error": err.Error()})
		case *customerrors.InternalServerErrorType:
			c.JSON(errCodeInternalServerError, gin.H{"error": err.Error()})
		case *customerrors.NotFoundErrorType:
			c.JSON(errCodeNotFound, gin.H{"error": err.Error()})
		case *customerrors.UnprocessableEntityErrorType:
			c.JSON(errCodeUnprocessableEntity, gin.H{"error": err.Error()})
		default:
			c.JSON(errCodeInternalServerError, gin.H{"error": err.Error()}) // 予期せぬエラーの場合、500を返す
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
