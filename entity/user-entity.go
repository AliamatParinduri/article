package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(50);not null" json:"name"`
	Username  string `gorm:"type:varchar(25);not null;uniqueIndex" json:"username"`
	Password  string `gorm:"type:varchar(100);not null" json:"password"`
	IsAdmin   bool   `gorm:"default:0" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
