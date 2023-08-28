package task

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/middleware"
	"github.com/vanntrong/asana-clone-be/project"
	"github.com/vanntrong/asana-clone-be/utils"
)

type TaskController struct {
	taskService    ITaskService
	projectService project.IProjectService
}

func registerRoutes(router *gin.RouterGroup, ctrl *TaskController) {
	v1 := router.Group("/tasks")
	v1.POST("/", middleware.AuthMiddleware, ctrl.Create)
	v1.GET("/:id", middleware.AuthMiddleware, ctrl.GetById)
	v1.PUT("/:id", middleware.AuthMiddleware, ctrl.UpdateTask)
	v1.DELETE("/:id", middleware.AuthMiddleware, ctrl.DeleteTask)
}

func NewTaskController(app *gin.RouterGroup, taskService ITaskService, projectService project.IProjectService) {
	ctrl := &TaskController{taskService, projectService}
	registerRoutes(app, ctrl)
}

func (ctrl *TaskController) Create(ctx *gin.Context) {
	var body CreateTaskValidation

	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	userId := ctx.GetHeader(configs.HeaderUserId)

	projectMember, err := ctrl.projectService.GetListMember(body.ProjectId)

	if err != nil || len(*projectMember) == 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("project not found"))
		return
	}

	if !project.IsMember(projectMember, userId) {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("you are not member of this project"))
		return
	}

	task, err := ctrl.taskService.Create(&body, userId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{
		"task": task.ID.String(),
	}, http.StatusOK)
}

func (ctrl *TaskController) GetById(ctx *gin.Context) {
	taskId := ctx.Param("id")
	userId := ctx.GetHeader(configs.HeaderUserId)
	task, err := ctrl.taskService.GetById(taskId, userId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{
		"task": task,
	}, http.StatusOK)
}

func (ctrl *TaskController) UpdateTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	userId := ctx.GetHeader(configs.HeaderUserId)

	var body UpdateTaskValidation

	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	task, err := ctrl.taskService.UpdateTask(taskId, &body, userId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{
		"task": task.ID.String(),
	}, http.StatusOK)
}

func (ctrl *TaskController) DeleteTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	userId := ctx.GetHeader(configs.HeaderUserId)

	err := ctrl.taskService.DeleteTask(taskId, userId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{}, http.StatusNoContent)
}
