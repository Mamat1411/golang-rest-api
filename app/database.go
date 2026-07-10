package app

import (
	"golang-rest-api/database/seeding"
	"golang-rest-api/models/domain"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dataSourceName string) {
	database, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if os.Getenv("DB_RESET") == "true" {
		err := database.Migrator().DropTable(
			&domain.Post{},
		)
		if err != nil {
			panic("Failed to drop table")
		}
	}

	database.AutoMigrate(
		&domain.Post{},
	)

	if os.Getenv("DB_RESET") == "true" {
		seeding.Seed(database)
	}

	DB = database
}
