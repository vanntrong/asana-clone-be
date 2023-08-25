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

	// init routes
	userRepository := user.NewUserRepository(db.DB)
	authService := auth.NewAuthService(userRepository)
	auth.NewAuthController(routes, authService)

	userService := user.NewUserService(userRepository)
	user.NewUserController(routes, userService)

	projectRepository := project.NewProjectRepository(db.DB)
	projectService := project.NewProjectService(projectRepository, userService)
	project.NewProjectController(routes, projectService)

	taskRepository := task.NewTaskRepository(db.DB)
	taskService := task.NewTaskService(taskRepository, projectService)
	task.NewTaskController(routes, taskService)
}
