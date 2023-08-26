package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(dbUrl string) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var err error
	log.Println(dbUrl)
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

}
