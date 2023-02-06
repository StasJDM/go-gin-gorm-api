package main

import (
	"github.com/StasJDM/go-gin-gorm-api/controllers"
	"github.com/StasJDM/go-gin-gorm-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	models.ConnectDB()

	route.GET("/users", controllers.FindUsers)
	route.GET("/users/:id", controllers.FindOneUser)
	route.POST("/users", controllers.CreateUser)
	route.PATCH("/users/:id", controllers.UpdateUser)
	route.DELETE("/users/:id", controllers.DeleteUser)

	route.Run()
}
