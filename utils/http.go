package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateResponse(ctx *gin.Context, data interface{}, code int, pagination ...interface{}) {

	if pagination != nil {
		ctx.JSON(code, gin.H{
			"message":    http.StatusText(code),
			"data":       data,
			"pagination": pagination,
		})
		return
	}
	ctx.JSON(code, gin.H{
		"message": http.StatusText(code),
		"data":    data,
	})
}
