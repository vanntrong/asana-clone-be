package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/utils"
)

type AuthController struct {
	authService IAuthService
}

func registerRoutes(router *gin.RouterGroup, ctrl *AuthController) {
	v1 := router.Group("/auth")
	v1.POST("/register", ctrl.RegisterUser)
	v1.POST("/login", ctrl.LoginUser)
	v1.POST("/check-email", ctrl.CheckEmail)
}

func NewAuthController(app *gin.RouterGroup, authService IAuthService) {
	ctrl := &AuthController{
		authService,
	}
	registerRoutes(app, ctrl)
}

func (ctrl *AuthController) RegisterUser(ctx *gin.Context) {
	var body RegisterValidation
	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	var _, token, err = ctrl.authService.Register(body)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, token, http.StatusOK)
}

func (ctrl *AuthController) LoginUser(ctx *gin.Context) {
	var body LoginValidation
	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	var _, token, err = ctrl.authService.Login(body)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid email or password"))
		return
	}

	utils.GenerateResponse(ctx, token, http.StatusOK)
}

func (ctrl *AuthController) CheckEmail(ctx *gin.Context) {
	var body CheckEmailValidation

	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	var info, err = ctrl.authService.CheckEmail(body)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Email is not exist"))
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{
		"email":  info.Email,
		"avatar": info.Avatar,
	}, http.StatusOK)
}
