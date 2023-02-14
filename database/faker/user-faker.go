package faker

import (
	"article_app/entity"
	"time"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *entity.User {
	return &entity.User{
		ID:        1,
		Name:      faker.FirstName(),
		Username:  faker.LastName(),
		Password:  "$2y$10$92ixunpkjo0roq5bymi.ye4okoea3ro9llc/.og/at2.uhewg/igi",
		IsAdmin:   true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}
