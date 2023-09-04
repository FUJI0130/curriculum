// FWからのリクエストをユースケースに送り、レスポンスを返す
package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
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

	log.Printf("[DEBUG] start Create create_user_controller.go")
	var req userapp.CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[DEBUG] before Exec Create create_user_controller.go")
	if err := ctrl.createUserService.Exec(c, &req); err != nil {
		// エラーの種類によってHTTPステータスコードを変更

		log.Printf("[DEBUG] into Exec error handling1 Create create_user_controller.go")
		if errors.Is(err, userapp.ErrUserNameAlreadyExists) {
			log.Printf("[DEBUG] into Exec error handling2 Create create_user_controller.go")
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			log.Printf("[DEBUG] into Exec error handling3 else  Create create_user_controller.go")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
