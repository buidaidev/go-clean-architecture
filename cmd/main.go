package main

import (
	"go-clean-architecture/app"
	"go-clean-architecture/internal/middleware"
	"go-clean-architecture/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	InitialEnvironment()

	db := mysql.NewMySqlDB()
	appContext := app.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	v1 := r.Group("/v1")
	setupRoutes(appContext, v1)

	r.Run(":8080")
}
