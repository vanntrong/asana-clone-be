package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseError struct {
	Error      string
	Message    string
	StatusCode int
}

func NewError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"error":      http.StatusText(code),
		"statusCode": code,
		"message":    message,
	})
}
