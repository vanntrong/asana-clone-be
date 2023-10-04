package migrations

import (
	"log"

	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/db"
	"github.com/vanntrong/asana-clone-be/entities"
)

func init() {
	err := configs.LoadEnv(".", &configs.AppConfig)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectToDatabase(configs.AppConfig.DBUrl)
}

func AutoMigrate() {
	MigrationTable()
	SetupJoinTable()
}

func SetupJoinTable() {
	db.DB.SetupJoinTable(&entities.Project{}, "Users", &entities.ProjectUsers{})
	db.DB.SetupJoinTable(&entities.Task{}, "Likes", &entities.TaskLikes{})
}

func MigrationTable() {
	db.DB.AutoMigrate(&entities.User{})
	db.DB.AutoMigrate(&entities.Project{})
	db.DB.AutoMigrate(&entities.ProjectUsers{})
	db.DB.AutoMigrate(&entities.TaskLikes{})
	db.DB.AutoMigrate(&entities.Task{})
	db.DB.AutoMigrate(&entities.Comment{})
	db.DB.AutoMigrate(&entities.Section{})
}
