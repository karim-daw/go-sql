package models

import (
	"errors"
	"time"
)

type Post struct {
	ID        uint      `json:"post_id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  uint      `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Post) Validate() error {

	if p.Title == "" {
		return errors.New("required title")
	}
	if p.Content == "" {
		return errors.New("required content")
	}
	return nil
}
