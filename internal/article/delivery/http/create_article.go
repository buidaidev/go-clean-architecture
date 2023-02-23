package articlehandler

import (
	"go-clean-architecture/domain"
	articlerepo "go-clean-architecture/internal/article/repository"
	articleus "go-clean-architecture/internal/article/usecase"

	"go-clean-architecture/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticle(appCtx app.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()

		var data domain.ArticleCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(app.ErrorInvalidRequest(err))
		}

		repo := articlerepo.NewArticleRepo(db)
		us := articleus.NewArticleUseCase(repo)
		ctx := c.Request.Context()

		if err := us.CreateArticle(ctx, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, app.SimpleSuccessRespone(data))
	}
}
