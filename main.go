package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/db"
	"github.com/vanntrong/asana-clone-be/migrations"
)

func init() {
	err := configs.LoadEnv(".", &configs.AppConfig)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectToDatabase(configs.AppConfig.DBUrl)
	migrations.AutoMigrate()
}

func main() {

	app := gin.Default()
	InitRoutes(app)
	address := ":" + configs.AppConfig.PORT
	app.Run(address) // listen and serve on 0.0.0.0:8080
}
