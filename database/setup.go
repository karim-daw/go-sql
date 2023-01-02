package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// get env
func hasEnvFile() error {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return err
	}
	log.Print("Found .env file")
	return nil
}

func ConnectDatabase() {

	// check if env file exists
	if hasEnvFile() != nil {
		return
	}

	PASSWORD := loadEnvVar("POSTGRES_PASSWORD")
	USERNAME := loadEnvVar("POSTGRES_USERNAME")
	DB_NAME := loadEnvVar("DATABASE_NAME")
	DB_PORT := loadEnvVar("DATABASE_PORT")
	DB_HOST := loadEnvVar("DATABASE_HOST")

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

	database.AutoMigrate(&Post{}) // register Post model
	database.AutoMigrate(&User{}) // register Post model

	DB = database
}

// returns the env variable given string key
func loadEnvVar(key string) string {
	envVariable, exists := os.LookupEnv(key)
	if !exists {
		fmt.Printf("Could not find %s \n", key)
	}
	return envVariable
}
