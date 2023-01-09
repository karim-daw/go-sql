package main

import (
	"go-sql/controllers"
	"go-sql/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)
	router.POST("/register", controllers.RegisterUser)

	router.GET("/posts", controllers.FindPosts)
	router.GET("/posts/:id", controllers.FindPost) // here!
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
	router.DELETE("/posts", controllers.DeleteAllPosts)
	router.Run("localhost:8080")
}
