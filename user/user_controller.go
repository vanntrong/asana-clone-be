package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/middleware"
	"github.com/vanntrong/asana-clone-be/utils"
)

type UserController struct {
	userService IUserService
}

func registerRoutes(router *gin.RouterGroup, ctrl *UserController) {
	v1 := router.Group("/users")
	v1.GET("/me", middleware.AuthMiddleware, ctrl.GetMe)
}

func NewUserController(app *gin.RouterGroup, userService IUserService) {
	ctrl := &UserController{userService}
	registerRoutes(app, ctrl)
}

func (ctrl *UserController) GetMe(ctx *gin.Context) {
	user_id := ctx.Request.Header.Get(configs.HeaderUserId)

	user, err := ctrl.userService.GetById(user_id)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("User not found"))
		return
	}

	utils.GenerateResponse(ctx, gin.H{
		"user": user.UserSerializer(),
	}, http.StatusOK)
}
