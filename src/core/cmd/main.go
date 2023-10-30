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
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", config.Env.DbUser, config.Env.DbPassword, config.Env.DbHost, config.Env.DbPort, config.Env.DbName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := rdbimpl.NewUserRepository()
	tagRepo := rdbimpl.NewTagRepository()
	existService := userdm.NewExistByNameDomainService(userRepo)
	createUserService := userapp.NewCreateUserAppService(userRepo, tagRepo, existService)
	updateUserService := userapp.NewUpdateUserAppService(userRepo, tagRepo)

	r := gin.Default()
	r.Use(middleware.ErrorHandler)
	r.Use(middleware.TransactionHandler(db))

	controllers.InitControllers(r, createUserService, updateUserService)
	log.Println("Starting server on port:", config.Env.AppPort)
	r.Run(":" + config.Env.AppPort)
}
