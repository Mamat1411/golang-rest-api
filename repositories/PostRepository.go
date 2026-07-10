package repositories

import (
	"golang-rest-api/models/domain"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (repository *PostRepository) GetPosts() ([]domain.Post, error) {
	var posts []domain.Post
	err := repository.DB.Find(&posts).Error
	return posts, err
}
