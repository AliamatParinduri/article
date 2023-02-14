package repository

import (
	"article_app/database/seed"
	"article_app/entity"
	"article_app/helper"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type ServerPG struct {
	DB *gorm.DB
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (s *ServerPG) InitialPostgres() *gorm.DB {
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
	s.DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalln(err)
	}

	return s.DB
}

func (s *ServerPG) dbMigrate() {
	var err error

	for _, model := range entity.RegisterModelPG() {
		if err = s.DB.Debug().AutoMigrate(model.Model); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully")
}

func (s *ServerPG) InitCommands() {
	s.InitialPostgres()

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				s.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seed.DBSeed(s.DB)
				if err != nil {
					log.Fatalln(err)
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
