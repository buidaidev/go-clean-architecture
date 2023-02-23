package articlehandler

import (
	"fmt"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
	articlerepo "go-clean-architecture/internal/article/repository"
	articleus "go-clean-architecture/internal/article/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchArticleus(appCtx app.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()

		var pagingData app.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(app.ErrorInvalidRequest(err))
		}
		pagingData.Fullfill()

		var filter domain.ArticleFilter

		if err := c.ShouldBind(&filter); err != nil {
			panic(app.ErrorInvalidRequest(err))
		}
		filter.Status = []int{1}
		fmt.Println(filter)

		repo := articlerepo.NewArticleRepo(db)
		us := articleus.NewArticleUseCase(repo)
		ctx := c.Request.Context()

		result, err := us.FetchArticleus(ctx, &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, app.SuccessRespone(result, pagingData, filter))
	}
}
