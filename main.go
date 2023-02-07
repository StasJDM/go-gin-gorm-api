package main

import (
	"log"

	"github.com/StasJDM/go-gin-gorm-api/controllers"
	"github.com/StasJDM/go-gin-gorm-api/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	models.ConnectDB()
	serveApplication()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load envs from .env file")
	}
}

func serveApplication() {
	route := gin.Default()

	route.GET("/users", controllers.FindUsers)
	route.GET("/users/:id", controllers.FindOneUser)
	route.POST("/users", controllers.CreateUser)
	route.PATCH("/users/:id", controllers.UpdateUser)
	route.DELETE("/users/:id", controllers.DeleteUser)

	route.POST("/register", controllers.Register)
	route.POST("/login", controllers.Login)

	route.Run(":8000")
	route.Run("Server running on port 8000")
}
