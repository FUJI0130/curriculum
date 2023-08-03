package main

import (
	"github.com/FUJI0130/curriculum/src/core/infra/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	controllers.InitControllers(r)

	r.Run(":8080")
}
