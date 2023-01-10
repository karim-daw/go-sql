package types

import "gorm.io/gorm"

// defines body schema for request "CreateBlog"
type PostInput struct {
	gorm.Model
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	AuthorID uint   `json:"author_id"`
}
