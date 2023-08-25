package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/configs"
	error "github.com/vanntrong/asana-clone-be/error"
	"github.com/vanntrong/asana-clone-be/utils"
)

func AuthMiddleware(ctx *gin.Context) {
	bearer_token := ctx.GetHeader("Authorization")

	if bearer_token == "" {
		error.NewError(ctx, http.StatusUnauthorized, "Unauthorized")
		ctx.Abort()
		return
	}

	token_str := bearer_token[7:]

	token, err := utils.ValidateToken(token_str)

	if token == nil || err != nil {
		error.NewError(ctx, http.StatusUnauthorized, "Unauthorized")
		ctx.Abort()
		return
	}

	ctx.Request.Header.Set(configs.HeaderUserId, fmt.Sprintf("%v", token["id"]))
	ctx.Request.Header.Set(configs.HeaderUserEmail, fmt.Sprintf("%v", token["email"]))
	ctx.Request.Header.Set(configs.HeaderUserName, fmt.Sprintf("%v", token["name"]))

	ctx.Next()
}
