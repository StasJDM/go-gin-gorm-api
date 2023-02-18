package controllers

import (
	"net/http"

	"github.com/StasJDM/go-gin-gorm-api/inputs"
	"github.com/StasJDM/go-gin-gorm-api/models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func FindUsers(context *gin.Context) {
	var pagination inputs.PaginationInput

	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var users []models.User
	models.DB.Limit(pagination.Limit()).Offset(pagination.Offset()).Find(&users)

	context.JSON(http.StatusOK, gin.H{"data": users})
}

func FindOneUser(context *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(context *gin.Context) {
	var input CreateUserInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username, Email: input.Email}
	models.DB.Create(&user)

	context.JSON(http.StatusCreated, gin.H{"data": user})
}

func UpdateUser(context *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var input UpdateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Model(&user).Updates(models.User{Username: input.Username, Email: input.Email})

	context.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(context *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&user)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
