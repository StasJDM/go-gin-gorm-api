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

// Find all users godoc
// @Summary Find all users
// @Schemes
// @Description Fetch all users (with pagination)
// @Tags users
// @Accept json
// @Produce json
// @Param pagination query inputs.PaginationInput true "User id"
// @Success 200 {array} models.User
// @Router /users [get]
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

// Find user by id godoc
// @Summary Find user by id
// @Schemes
// @Description Find one user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func FindOneUser(context *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}

// Create user godoc
// @Summary Create user
// @Schemes
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserInput true "User json"
// @Success 200 {object} models.User
// @Router /users [post]
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

// Update user godoc
// @Summary Update user
// @Schemes
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Param user body UpdateUserInput true "User json"
// @Success 200 {object} models.User
// @Router /users/{id} [patch]
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

// Delete user godoc
// @Summary Delete user
// @Schemes
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Success 200
// @Router /users/{id} [delete]
func DeleteUser(context *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&user)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
