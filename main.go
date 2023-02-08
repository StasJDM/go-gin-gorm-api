package main

import (
	"log"

	"github.com/StasJDM/go-gin-gorm-api/controllers"
	"github.com/StasJDM/go-gin-gorm-api/middlewares"
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
	router := gin.Default()

	userRoutes := router.Group("/users")
	userRoutes.Use(middlewares.JWTAuthMiddleware())
	userRoutes.GET("", controllers.FindUsers)
	userRoutes.GET("/:id", controllers.FindOneUser)
	userRoutes.POST("", controllers.CreateUser)
	userRoutes.PATCH("/:id", controllers.UpdateUser)
	userRoutes.DELETE("/:id", controllers.DeleteUser)

	authRoutes := router.Group("/auth")
	authRoutes.POST("/register", controllers.Register)
	authRoutes.POST("/login", controllers.Login)

	router.Run(":8000")
	router.Run("Server running on port 8000")
}
