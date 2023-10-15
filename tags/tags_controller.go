package tags

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/middleware"
	"github.com/vanntrong/asana-clone-be/utils"
)

type TagsController struct {
	tagsService ITagsService
}

func registerRoutes(router *gin.RouterGroup, ctrl *TagsController) {
	v1 := router.Group("/tags")
	v1.GET("/", middleware.AuthMiddleware, ctrl.GetList)
	v1.POST("/", middleware.AuthMiddleware, ctrl.Create)
	v1.PUT("/:id", middleware.AuthMiddleware, ctrl.Update)
}

func NewTagsController(app *gin.RouterGroup, tagsService ITagsService) {
	ctrl := &TagsController{tagsService}
	registerRoutes(app, ctrl)
}

func (ctrl *TagsController) GetList(ctx *gin.Context) {
	var query GetListTagsValidation

	isValid := utils.ValidationQuery(ctx, &query)

	if !isValid {
		return
	}

	userId := ctx.GetHeader(configs.HeaderUserId)

	tags, err := ctrl.tagsService.FindTags(userId, query.ProjectId)

	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	utils.GenerateResponse(ctx, tags, http.StatusOK)
}

func (ctrl *TagsController) Create(ctx *gin.Context) {
	var body CreateTagValidation

	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	userId := ctx.GetHeader(configs.HeaderUserId)

	tag, err := ctrl.tagsService.CreateTag(userId, body)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, tag, http.StatusCreated)
}

func (ctrl *TagsController) Update(ctx *gin.Context) {
	var body UpdateTagValidation

	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	userId := ctx.GetHeader(configs.HeaderUserId)

	tagId := ctx.Param("id")

	tag, err := ctrl.tagsService.UpdateTag(tagId, userId, body)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, tag, http.StatusOK)
}
