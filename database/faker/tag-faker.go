package faker

import (
	"article_app/entity"
	"time"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func TagFaker(db *gorm.DB) *entity.Tag {
	for i := 0; i <= 5; i++ {
		db.Create(&entity.Tag{
			ID:        i,
			Name:      faker.Username(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		})
	}
	return nil
}
