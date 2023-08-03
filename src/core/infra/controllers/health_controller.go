package controllers

import "github.com/gin-gonic/gin"

func InitHealthController(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
