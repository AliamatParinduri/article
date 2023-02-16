package entity

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	ID        int    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(25);not null" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
