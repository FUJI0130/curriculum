package controllers

import (
	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/gin-gonic/gin"
)

func InitControllers(r *gin.Engine, createUserService *userapp.CreateUserAppService) {
	InitHealthController(r)
	InitCreateUserController(r, createUserService)
}

func InitCreateUserController(r *gin.Engine, s *userapp.CreateUserAppService) {
	ctrl := NewCreateUserController(s)
	r.POST("/users", ctrl.Create)
}

func InitUpdateUserController(r *gin.Engine, UpdateuserService *userapp.UpdateUserAppService) {
	ctrl := NewUpdateUserController(UpdateuserService)
	r.GET("/users/:userID", ctrl.Fetch)
	r.PATCH("/users", ctrl.Update)
}
