package controllers

import (
	"go-sql/database"
	"go-sql/models"
)

// TODO:inject #1 model dependancy in controller
func Migrate() {
	database.DB.AutoMigrate(&models.Post{}) // register Post model
	database.DB.AutoMigrate(&models.User{}) // register User model
}
