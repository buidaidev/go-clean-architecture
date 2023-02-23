package middleware

import (
	"go-clean-architecture/app"

	"github.com/gin-gonic/gin"
)

func Recover(appCtx app.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if appError, ok := err.(*app.AppError); ok {
					ctx.AbortWithStatusJSON(appError.StatusCode, appError)
					return
				}

				appError := app.ErrorInternal(err.(error))
				ctx.AbortWithStatusJSON(appError.StatusCode, appError)
				return
			}
		}()

		ctx.Next()
	}
}
