package controllers

import (
	"go-sql/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// defines body schema for request "CreateBlog"
type CreatePostInput struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	AuthorID uint   `json:"author_id"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	log.Print("creating post")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.Post{Title: input.Title, Content: input.Content, AuthorID: input.AuthorID}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// find every post entry in database
func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func FindPost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})

}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id =?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var input UpdatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedPost := models.Post{Title: input.Title, Content: input.Content}

	models.DB.Model(&post).Updates(&updatedPost)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	models.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}

func DeleteAllPosts(c *gin.Context) {
	var posts []models.Post
	result := models.DB.Find(&posts).RowsAffected

	if result == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "no records found"})
		return
	}
	models.DB.Delete(&posts)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}

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
	user.SaveUser()

	c.JSON(http.StatusOK, gin.H{"data": user})
}
