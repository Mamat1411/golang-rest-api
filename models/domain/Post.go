package domain

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	Id        string         `json:"id" gorm:"primary_key"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
