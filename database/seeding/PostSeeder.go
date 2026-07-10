package seeding

import (
	"golang-rest-api/models/domain"
	"log"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

func PostSeeder(db *gorm.DB) {
	var posts []domain.Post

	for range 10 {
		posts = append(posts, domain.Post{
			Id:      gofakeit.UUID(),
			Title:   gofakeit.BookTitle(),
			Content: gofakeit.Paragraph(),
		})
	}

	if err := db.Create(&posts).Error; err != nil {
		log.Fatalf("Failed to seed post: %v", err)
	}

	log.Println("Post Seeding Success")
}
