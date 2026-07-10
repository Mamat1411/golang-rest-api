package domain

type Post struct {
	Id      string `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
