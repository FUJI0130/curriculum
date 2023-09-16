package main

import (
	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm/interfaces"
	"github.com/FUJI0130/curriculum/src/core/infra/controllers"
	"github.com/FUJI0130/curriculum/src/core/infra/rdbimpl"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	// DB接続の設定
	db, err := sqlx.Open("mysql", "user:password@tcp(mysql:3306)/sql?parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := rdbimpl.NewUserRepository(db)
	tagRepo := rdbimpl.NewTagRepository(db)
	existService := interfaces.NewExistByNameDomainService(userRepo)
	createUserService := userapp.NewCreateUserAppService(userRepo, tagRepo, existService)

	r := gin.Default()
	controllers.InitControllers(r, createUserService)
	r.Run(":8080")
}
