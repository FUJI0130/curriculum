// src/core/cmd/main.go

package main

import (
	"fmt"

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
	env := config.LoadEnv()

	// DB接続の設定
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", env.DbUser, env.DbPassword, env.DbHost, env.DbPort, env.DbName))
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
	r.Use(middleware.TransactionHandler(db))
	controllers.InitControllers(r, createUserService)
	r.Run(":" + env.AppPort)
}
