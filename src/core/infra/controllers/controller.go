package controllers

import "github.com/gin-gonic/gin"

func InitControllers(r *gin.Engine) {
	InitHealthController(r)
	// Here you can add other controllers initializers
	// InitUserController(r)
	// InitSomeOtherController(r)
}
