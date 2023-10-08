package main

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/auth"
	"github.com/vanntrong/asana-clone-be/comments"
	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/db"
	"github.com/vanntrong/asana-clone-be/middleware"
	"github.com/vanntrong/asana-clone-be/project"
	"github.com/vanntrong/asana-clone-be/sections"
	"github.com/vanntrong/asana-clone-be/task"
	"github.com/vanntrong/asana-clone-be/user"
)

func InitRoutes(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	originList := strings.Split(configs.AppConfig.OriginHost, ",")
	fmt.Println(originList)
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-CSRF-Token", "Access-Control-Allow-Origin"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowOriginFunc = func(origin string) bool {
		for _, v := range originList {
			fmt.Println("v:", v)
			fmt.Println("origin:", origin)
			return v == origin
		}
		return false
	}

	app.Use(cors.New(config))

	app.Use(middleware.ErrorHandler)
	routes := app.Group("/api/v1")
	// init repository
	userRepository := user.NewUserRepository(db.DB)
	projectRepository := project.NewProjectRepository(db.DB)
	taskRepository := task.NewTaskRepository(db.DB)
	sectionsRepository := sections.NewSectionsRepository(db.DB)
	commentsRepository := comments.NewCommentsRepository(db.DB)

	// init service and controller
	authService := auth.NewAuthService(userRepository)
	userService := user.NewUserService(userRepository)
	projectService := project.NewProjectService(projectRepository, userService)
	taskService := task.NewTaskService(taskRepository, projectService)
	sectionsService := sections.NewSectionsService(sectionsRepository, projectService)
	commentsService := comments.NewCommentsService(commentsRepository, taskService, projectService)

	auth.NewAuthController(routes, authService)
	user.NewUserController(routes, userService)
	project.NewProjectController(routes, projectService)
	task.NewTaskController(routes, taskService, projectService)
	sections.NewSectionsController(routes, sectionsService)
	comments.NewCommentsController(routes, commentsService)
}
