// FWからのリクエストをユースケースに送り、レスポンスを返す
package controllers

import (
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	databaseErrors "github.com/FUJI0130/curriculum/src/core/common/database_errors"
	"github.com/FUJI0130/curriculum/src/core/common/errorcodes"
	domainErrors "github.com/FUJI0130/curriculum/src/core/domain/customerrors"
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
		c.JSON(errorcodes.BadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.createUserService.Exec(c, &req); err != nil {
		switch err.(type) {
		case *domainErrors.UserNameAlreadyExistsError:
			c.JSON(errorcodes.Conflict, gin.H{"error": err.Error()})
		case *domainErrors.TagNameAlreadyExistsError:
			c.JSON(errorcodes.BadRequest, gin.H{"error": err.Error()})
		case *domainErrors.DuplicateSkillTagError:
			c.JSON(errorcodes.BadRequest, gin.H{"error": err.Error()})
		case *databaseErrors.DatabaseError:
			c.JSON(errorcodes.InternalServerError, gin.H{"error": err.Error()})
		default:
			c.JSON(errorcodes.InternalServerError, gin.H{"error": err.Error()}) // 予期せぬエラーの場合、500を返す
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
