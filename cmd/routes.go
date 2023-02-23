package main

import (
	"go-clean-architecture/app"
	articlehandler "go-clean-architecture/internal/article/delivery/http"

	"github.com/gin-gonic/gin"
)

func setupRoutes(appCtx app.AppContext, r *gin.RouterGroup) {
	// article
	articles := r.Group("articles")
	{
		articles.POST("/", articlehandler.CreateArticle(appCtx))
		articles.GET("/", articlehandler.FetchArticleus(appCtx))
		articles.GET("/:id", articlehandler.GetArticleus(appCtx))
		articles.PATCH("/:id", articlehandler.UpdateArticleus(appCtx))
		articles.DELETE("/:id", articlehandler.DeleteArticleus(appCtx))
	}
}
