package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateResponse(ctx *gin.Context, data map[string]interface{}, code int) {

	ctx.JSON(code, gin.H{
		"message": http.StatusText(code),
		"data":    data,
	})
}
