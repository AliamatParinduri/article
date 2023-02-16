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
		Password:  "$2a$10$GAPngBcZY5xouanHXSYQCuFcT8ANkG1DhmCoFHa.E0rUvQ30Vw51S", // password
		IsAdmin:   true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}
