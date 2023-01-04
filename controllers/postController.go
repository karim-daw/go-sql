package controllers

import (
	"go-sql/database"
	"go-sql/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// defines body schema for request "CreateBlog"
type CreatePostInput struct {
	gorm.Model
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
	database.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// find every post entry in database
func FindPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// finds a post based on post ID
func FindPost(c *gin.Context) {
	var post models.Post
	if err := database.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// struct representing updated post
type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// updates a post
func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := database.DB.Where("id =?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedPost := models.Post{Title: input.Title, Content: input.Content}

	database.DB.Model(&post).Updates(&updatedPost)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := database.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	database.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}

func DeleteAllPosts(c *gin.Context) {
	var posts []models.Post
	result := database.DB.Find(&posts).RowsAffected

	if result == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "no records found"})
		return
	}
	database.DB.Delete(&posts)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
