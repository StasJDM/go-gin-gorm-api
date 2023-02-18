package controllers

import (
	"net/http"

	"github.com/StasJDM/go-gin-gorm-api/helpers"
	"github.com/StasJDM/go-gin-gorm-api/inputs"
	"github.com/StasJDM/go-gin-gorm-api/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(context *gin.Context) {
	var input inputs.CreatePostInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserId:  user.ID,
	}
	models.DB.Create(&post)

	context.JSON(http.StatusCreated, gin.H{"data": post})
}

func FindPosts(context *gin.Context) {
	var pagination inputs.PaginationInput

	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var posts []models.Post
	models.DB.Limit(pagination.Limit()).Offset(pagination.Offset()).Find(&posts)

	context.JSON(http.StatusOK, gin.H{"data": posts})
}

func FindOnePost(context *gin.Context) {
	var post models.Post

	if err := models.DB.Where("id = ?", context.Param("id")).First(&post).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
	}

	context.JSON(http.StatusOK, gin.H{"data": post})
}
