package router

import (
	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/infra/controllers"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine, createUserService *userapp.CreateUserAppService, updateUserService *userapp.UpdateUserAppService) {
	InitCreateUserController(r, createUserService)
	InitUpdateUserController(r, updateUserService)
}

func InitCreateUserController(r *gin.Engine, s *userapp.CreateUserAppService) {
	ctrl := controllers.NewCreateUserController(s)
	r.POST("/users", ctrl.Create)
}

func InitUpdateUserController(r *gin.Engine, UpdateuserService *userapp.UpdateUserAppService) {
	ctrl := controllers.NewUpdateUserController(UpdateuserService)
	r.PATCH("/users", ctrl.Update)
}
