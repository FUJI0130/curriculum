// FWからのリクエストをユースケースに送り、レスポンスを返す
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

// ここでcurlコマンドの内容をバインドしている
func (ctrl *CreateUserController) Create(c *gin.Context) {

	var req userapp.CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(customerrors.ErrCodeUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.createUserService.Exec(c, &req); err != nil {
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
