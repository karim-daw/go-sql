package models

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
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var PASSWORD = loadEnvVar("POSTGRES_PASSWORD")
var USERNAME = loadEnvVar("POSTGRES_USERNAME")
var DB_NAME = loadEnvVar("DATABASE_NAME")
var DB_PORT = loadEnvVar("DATABASE_PORT")
var DB_HOST = loadEnvVar("DATABASE_HOST")

func ConnectDatabase() {

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

	DB = database
}

func loadEnvVar(v string) string {
	envVariable, exists := os.LookupEnv(v)
	if exists {
		fmt.Printf("Found %s", v)
	}
	return envVariable
}
