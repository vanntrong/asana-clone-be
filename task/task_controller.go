package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/middleware"
	"github.com/vanntrong/asana-clone-be/utils"
)

type TaskController struct {
	taskService ITaskService
}

func registerRoutes(router *gin.RouterGroup, ctrl *TaskController) {
	v1 := router.Group("/tasks")
	v1.POST("/", middleware.AuthMiddleware, ctrl.Create)
	v1.GET("/:id", middleware.AuthMiddleware, ctrl.GetById)
}

func NewTaskController(app *gin.RouterGroup, taskService ITaskService) {
	ctrl := &TaskController{taskService}
	registerRoutes(app, ctrl)
}

func (ctrl *TaskController) Create(ctx *gin.Context) {
	var body CreateTaskValidation

	isValid := utils.Validation(ctx, &body)

	if !isValid {
		return
	}

	userId := ctx.GetHeader(configs.HeaderUserId)

	task, err := ctrl.taskService.Create(&body, userId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	utils.GenerateResponse(ctx, map[string]interface{}{
		"task": task,
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
