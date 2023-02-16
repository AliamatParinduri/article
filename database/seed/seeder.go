package seed

import (
	"article_app/database/faker"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeedersPG(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: faker.UserFaker(db)},
		{Seeder: faker.TagFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeedersPG(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}
	return nil
}
