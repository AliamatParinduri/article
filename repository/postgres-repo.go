package repository

import (
	"article_app/entity"
	"article_app/helper"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func InitialPostgres() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	var DBConfig = DBConfig{
		DBHost:     helper.GetEnv("DB_HOST", "localhost"),
		DBUser:     helper.GetEnv("DB_USER", "postgres"),
		DBPassword: helper.GetEnv("DB_PASSWORD", "password"),
		DBName:     helper.GetEnv("DB_NAME", "dbname"),
		DBPort:     helper.GetEnv("DB_PORT", "5432"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timeZone=Asia/Jakarta", DBConfig.DBHost, DBConfig.DBUser, DBConfig.DBPassword, DBConfig.DBName, DBConfig.DBPort)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal(err)
	}
	dbMigrate(db)

	return db
}

func dbMigrate(db *gorm.DB) {
	var err error

	for _, model := range entity.RegisterModelPG() {
		if err = db.Debug().AutoMigrate(model.Model); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully")
}
