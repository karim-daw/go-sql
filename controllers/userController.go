package controllers

import (
	"go-sql/database"
	"go-sql/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	log.Print("creating user")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Username: input.Username}
	user.SaveUser() // TODO - handle error

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUser(c *gin.Context) {
	var user models.User
	database.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
}

func DeleteUser(c *gin.Context) {
}
