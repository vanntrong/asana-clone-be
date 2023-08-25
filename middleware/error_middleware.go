package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {

		err := ctx.Errors.Last()
		errMessage := err.Err.Error()
		errStatusCode := ctx.Writer.Status()
		ctx.JSON(errStatusCode, gin.H{
			"error":      http.StatusText(errStatusCode),
			"statusCode": errStatusCode,
			"message":    errMessage,
		})

	}
}
