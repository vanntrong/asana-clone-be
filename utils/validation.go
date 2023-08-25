package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validation(ctx *gin.Context, obj interface{}) bool {
	validate := validator.New()

	fmt.Println(obj)

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
