package main

import (
	"fmt"
	"log"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/config"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/infra/controllers"
	"github.com/FUJI0130/curriculum/src/core/infra/middleware"
	"github.com/FUJI0130/curriculum/src/core/infra/rdbimpl"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func myHandler(c *gin.Context) {
	panic("this is a test panic")
}

func myBaseErrorHandler(c *gin.Context) {
	panic(customerrors.NewBaseError("This is a BaseError panic", 404, nil))
}

func main() {

	config.GlobalConfig.DebugMode = true

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", config.Env.DbUser, config.Env.DbPassword, config.Env.DbHost, config.Env.DbPort, config.Env.DbName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := rdbimpl.NewUserRepository(db)
	tagRepo := rdbimpl.NewTagRepository(db)
	existService := userdm.NewExistByNameDomainService(userRepo)
	createUserService := userapp.NewCreateUserAppService(userRepo, tagRepo, existService)

	r := gin.Default()
	r.Use(middleware.ErrorHandler)

	r.GET("/panic", myHandler)
	r.GET("/base-error-panic", myBaseErrorHandler)

	controllers.InitControllers(r, createUserService)
	log.Println("Starting server on port:", config.Env.AppPort)
	r.Run(":" + config.Env.AppPort)
}
