package main

import (
	"../infra/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	controllers.InitControllers(r)

	r.Run(":8080")
}
