package app

import (
	"golang-rest-api/models/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dataSourceName string) {
	database, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&domain.Post{})

	DB = database
}
