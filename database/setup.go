package database

import (
	"fmt"
	"log"

	"go-sql/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// check if env file exists
	if helpers.HasEnvFile() != nil {
		log.Print(helpers.HasEnvFile().Error())
		return
	}

	PASSWORD := helpers.LoadEnvVar("POSTGRES_PASSWORD")
	USERNAME := helpers.LoadEnvVar("POSTGRES_USERNAME")
	DB_NAME := helpers.LoadEnvVar("DATABASE_NAME")
	DB_PORT := helpers.LoadEnvVar("DATABASE_PORT")
	DB_HOST := helpers.LoadEnvVar("DATABASE_HOST")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timezone=Asia/Shanghai",
		DB_HOST,
		USERNAME,
		PASSWORD,
		DB_NAME,
		DB_PORT)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // change the database provider if necessary

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
