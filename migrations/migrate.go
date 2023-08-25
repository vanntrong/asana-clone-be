package migrations

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vanntrong/asana-clone-be/db"
	"github.com/vanntrong/asana-clone-be/entities"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectToDatabase()
}

func AutoMigrate() {
	MigrationTable()
	SetupJoinTable()
}

func SetupJoinTable() {
	db.DB.SetupJoinTable(&entities.Project{}, "Users", &entities.ProjectUsers{})
}

func MigrationTable() {
	db.DB.AutoMigrate(&entities.User{})
	db.DB.AutoMigrate(&entities.Project{})
	db.DB.AutoMigrate(&entities.ProjectUsers{})
	db.DB.AutoMigrate(&entities.Task{})
	db.DB.AutoMigrate(&entities.Comment{})
}
