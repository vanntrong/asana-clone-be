package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vanntrong/asana-clone-be/db"
	"github.com/vanntrong/asana-clone-be/migrations"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectToDatabase()
	migrations.AutoMigrate()
}

func main() {
	app := gin.Default()
	InitRoutes(app)
	app.Run() // listen and serve on 0.0.0.0:8080
}
