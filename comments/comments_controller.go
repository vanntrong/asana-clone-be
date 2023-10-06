package comments

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/middleware"
	"github.com/vanntrong/asana-clone-be/utils"
)

type CommentsController struct {
	commentsService ICommentsService
}

func NewCommentsController(app *gin.RouterGroup, commentsService ICommentsService) {
	ctrl := &CommentsController{commentsService}
	registerRoutes(app, ctrl)
}

func registerRoutes(router *gin.RouterGroup, ctrl *CommentsController) {
	v1 := router.Group("/comments")
	v1.GET("/", middleware.AuthMiddleware, ctrl.GetList)
	v1.POST("/", middleware.AuthMiddleware, ctrl.Create)
	v1.PATCH("/:id", middleware.AuthMiddleware, ctrl.Update)
	v1.DELETE("/:id", ctrl.Delete)
}

func (ctrl *CommentsController) Create(ctx *gin.Context) {
	var body CreateCommentValidation
	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	authorId := ctx.GetHeader(configs.HeaderUserId)
	comment, err := ctrl.commentsService.Create(authorId, body)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, comment, http.StatusOK)
}

func (ctrl *CommentsController) GetList(ctx *gin.Context) {
	var query GetListByTaskIdValidation
	isValid := utils.ValidationQuery(ctx, &query)

	if !isValid {
		return
	}

	userId := ctx.GetHeader(configs.HeaderUserId)
	comments, pagination, err := ctrl.commentsService.GetListByTaskId(userId, query)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, comments, http.StatusOK, pagination)
}

func (ctrl *CommentsController) Update(ctx *gin.Context) {
	var body UpdateCommentValidation
	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}
	userId := ctx.GetHeader(configs.HeaderUserId)
	commentId := ctx.Param("id")

	_, err := ctrl.commentsService.Update(userId, commentId, body)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{}, http.StatusNoContent)
}

func (ctrl *CommentsController) Delete(ctx *gin.Context) {
	userId := ctx.GetHeader(configs.HeaderUserId)
	commentId := ctx.Param("id")

	err := ctrl.commentsService.Delete(userId, commentId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{}, http.StatusNoContent)
}
