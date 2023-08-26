package project

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/configs"
)

func (ctrl *ProjectController) IsMemberOfProject(ctx *gin.Context) {
	projectId := ctx.Param("id")
	projectMembers, err := ctrl.projectService.GetListMember(projectId)
	userId := ctx.GetHeader(configs.HeaderUserId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if !IsMember(projectMembers, userId) {
		ctx.AbortWithError(http.StatusNotFound, errors.New("project not found"))
		return
	}

	ctx.Next()
}

func (ctrl *ProjectController) IsManagerOfProject(ctx *gin.Context) {
	projectId := ctx.Param("id")
	userId := ctx.GetHeader(configs.HeaderUserId)
	projectMembers, err := ctrl.projectService.GetListMember(projectId)

	if err != nil || len(*projectMembers) == 0 {
		ctx.AbortWithError(http.StatusNotFound, errors.New("project not found"))
		return
	}

	if !IsManager(projectMembers, userId) {
		ctx.AbortWithError(http.StatusForbidden, errors.New("you are not manager of this project"))
		return
	}
}
