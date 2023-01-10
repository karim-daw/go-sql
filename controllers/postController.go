package controllers

import (
	"go-sql/database"
	"go-sql/models"
	"go-sql/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create post
func CreatePost(c *gin.Context) {
	var input types.PostInput
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

// updates a post
func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := database.DB.Where("id =?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	var input types.UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedPost := models.Post{Title: input.Title, Content: input.Content}

	database.DB.Model(&post).Updates(&updatedPost)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// delete post
func DeletePost(c *gin.Context) {
	var post models.Post
	if err := database.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	database.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}

// delete all posts
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
