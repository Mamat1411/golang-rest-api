package services

import (
	"golang-rest-api/models/domain"
	"golang-rest-api/repositories"
)

type PostService struct {
	Repository *repositories.PostRepository
}

func NewPostService(repository *repositories.PostRepository) *PostService {
	return &PostService{Repository: repository}
}

func (service *PostService) GetPosts() ([]domain.Post, error) {
	return service.Repository.GetPosts()
}
