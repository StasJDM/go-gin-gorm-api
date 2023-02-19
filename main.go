package main

import (
	"log"

	"github.com/StasJDM/go-gin-gorm-api/controllers"
	_ "github.com/StasJDM/go-gin-gorm-api/docs"
	"github.com/StasJDM/go-gin-gorm-api/middlewares"
	"github.com/StasJDM/go-gin-gorm-api/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger API
// @version         1.0

// @host      localhost:8000

// @securityDefinitions.basic  BasicAuth
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

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRoutes := router.Group("/users")
	userRoutes.Use(middlewares.JWTAuthMiddleware())
	userRoutes.GET("", controllers.FindUsers)
	userRoutes.GET("/:id", controllers.FindOneUser)
	userRoutes.POST("", controllers.CreateUser)
	userRoutes.PATCH("/:id", controllers.UpdateUser)
	userRoutes.DELETE("/:id", controllers.DeleteUser)

	postRoutes := router.Group("/posts")
	postRoutes.Use(middlewares.JWTAuthMiddleware())
	postRoutes.GET("", controllers.FindPosts)
	postRoutes.GET("/:id", controllers.FindOnePost)
	postRoutes.POST("", controllers.CreatePost)

	authRoutes := router.Group("/auth")
	authRoutes.POST("/register", controllers.Register)
	authRoutes.POST("/login", controllers.Login)

	router.Run(":8000")
	router.Run("Server running on port 8000")
}
