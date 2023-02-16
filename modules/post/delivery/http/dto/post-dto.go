package dto

import "article_app/entity"

type PostDTO struct {
	UserID int          `gorm:"index" json:"user_id"`
	Tags   []entity.Tag `json:"tags"`
	Title  string       `validate:"required,min=3,max=50" json:"title"`
	Body   string       `validate:"required,min=5" json:"body"`
}
