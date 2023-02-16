package entity

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        int `gorm:"primary_key;auto_increment" json:"id"`
	User      User
	UserID    int    `gorm:"index" json:"-"`
	Tags      []Tag  `gorm:"many2many:post_tags;"`
	Title     string `gorm:"type:varchar(50);not null" json:"title"`
	Body      string `gorm:"type:text;not null" json:"body"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
