package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validation(ctx *gin.Context, obj interface{}) bool {
	validate := validator.New()

	if err := ctx.ShouldBindJSON(obj); err != nil {

		ctx.AbortWithError(http.StatusBadRequest, err)
		return false
	}

	if err := validate.Struct(obj); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return false
	}

	return true
}

func ValidationQuery(ctx *gin.Context, obj interface{}) bool {
	validate := validator.New()

	if err := ctx.ShouldBindQuery(obj); err != nil {

		ctx.AbortWithError(http.StatusBadRequest, err)
		return false
	}

	if err := validate.Struct(obj); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return false
	}

	return true
}
