package articlehandler

import (
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
	articlerepo "go-clean-architecture/internal/article/repository"
	articleus "go-clean-architecture/internal/article/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateArticleus(appCtx app.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(app.ErrorInvalidRequest(err))
		}

		var data domain.ArticleUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(app.ErrorInvalidRequest(err))
		}

		repo := articlerepo.NewArticleRepo(db)
		us := articleus.NewArticleUseCase(repo)
		ctx := c.Request.Context()

		if err := us.UpdateArticleus(ctx, &data, id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, app.SimpleSuccessRespone(data))
	}
}
