package articlehandler

import (
	"go-clean-architecture/app"
	articlerepo "go-clean-architecture/internal/article/repository"
	articleus "go-clean-architecture/internal/article/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteArticleus(appCtx app.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(app.ErrorInvalidRequest(err))
		}

		repo := articlerepo.NewArticleRepo(db)
		us := articleus.NewArticleUseCase(repo)
		ctx := c.Request.Context()

		if err := us.DeleteArticleus(ctx, id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, app.SimpleSuccessRespone("Delete successfully"))
	}
}
