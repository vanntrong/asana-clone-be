package project

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/middleware"
	"github.com/vanntrong/asana-clone-be/utils"
)

type ProjectController struct {
	projectService IProjectService
}

func registerRoutes(router *gin.RouterGroup, ctrl *ProjectController) {
	v1 := router.Group("/projects")
	v1.POST("/", middleware.AuthMiddleware, ctrl.Create)
	v1.GET("/:id", middleware.AuthMiddleware, ctrl.IsMemberOfProject, ctrl.GetById)
	v1.PATCH("/:id/members/add", middleware.AuthMiddleware, ctrl.IsManagerOfProject, ctrl.AddMember)
	v1.PATCH("/:id/members/remove", middleware.AuthMiddleware, ctrl.IsManagerOfProject, ctrl.RemoveMember)
}

func NewProjectController(app *gin.RouterGroup, projectService IProjectService) {
	ctrl := &ProjectController{projectService}
	registerRoutes(app, ctrl)
}

func (ctrl *ProjectController) Create(ctx *gin.Context) {
	var body CreateProjectValidation
	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	authorId := ctx.GetHeader(configs.HeaderUserId)

	project, err := ctrl.projectService.Create(body, authorId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{
		"project": project,
	}, http.StatusOK)

}

func (ctrl *ProjectController) GetById(ctx *gin.Context) {
	projectId := ctx.Param("id")

	project, err := ctrl.projectService.GetById(projectId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{
		"project": project,
	}, http.StatusOK)
}

func (ctrl *ProjectController) AddMember(ctx *gin.Context) {
	var body AddMemberValidation
	projectId := ctx.Param("id")
	userId := ctx.GetHeader(configs.HeaderUserId)

	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	if ctrl.isAddOrRemoveSelf(body.Members, userId) {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("you can not add yourself"))
		return
	}

	err := ctrl.projectService.AddMember(projectId, body, userId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{}, http.StatusNoContent)
}

func (ctrl *ProjectController) RemoveMember(ctx *gin.Context) {
	var body RemoveMemberValidation
	projectId := ctx.Param("id")
	userId := ctx.GetHeader(configs.HeaderUserId)

	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	if ctrl.isAddOrRemoveSelf(body.Members, userId) {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("you can not remove yourself"))
		return
	}

	err := ctrl.projectService.RemoveMember(projectId, body, userId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{}, http.StatusNoContent)
}

func (ctrl *ProjectController) isAddOrRemoveSelf(list []string, requestId string) bool {
	for _, id := range list {
		if id == requestId {
			return true
		}
	}

	return false
}
