package seeding

import (
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	log.Println("Database Seeding....")
	PostSeeder(db)
	log.Println("Seeding Completed")
}
