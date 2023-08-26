package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/auth"
	"github.com/vanntrong/asana-clone-be/db"
	"github.com/vanntrong/asana-clone-be/middleware"
	"github.com/vanntrong/asana-clone-be/project"
	"github.com/vanntrong/asana-clone-be/task"
	"github.com/vanntrong/asana-clone-be/user"
)

func InitRoutes(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	app.Use(middleware.ErrorHandler)

	routes := app.Group("/api/v1")

	// init repository
	userRepository := user.NewUserRepository(db.DB)
	projectRepository := project.NewProjectRepository(db.DB)
	taskRepository := task.NewTaskRepository(db.DB)

	// init service and controller
	authService := auth.NewAuthService(userRepository)
	userService := user.NewUserService(userRepository)
	projectService := project.NewProjectService(projectRepository, userService)
	taskService := task.NewTaskService(taskRepository, projectService)

	auth.NewAuthController(routes, authService)
	user.NewUserController(routes, userService)
	project.NewProjectController(routes, projectService)
	task.NewTaskController(routes, taskService)
}
