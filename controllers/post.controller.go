package controllers

import (
	"net/http"
	"time"

	"github.com/StasJDM/go-gin-gorm-api/cache"
	"github.com/StasJDM/go-gin-gorm-api/helpers"
	"github.com/StasJDM/go-gin-gorm-api/inputs"
	"github.com/StasJDM/go-gin-gorm-api/models"
	"github.com/gin-gonic/gin"
)

const POST_CACHE_PREFIX = "POST_"

// Create post godoc
// @Summary Create post
// @Schemes
// @Description Create post
// @Tags posts
// @Accept json
// @Produce json
// @Param user body inputs.CreatePostInput true "Post json"
// @Success 200 {object} models.Post
// @Router /posts [post]
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

// Find all posts godoc
// @Summary Find all posts
// @Schemes
// @Description Fetch all posts (with pagination)
// @Tags posts
// @Accept json
// @Produce json
// @Param pagination query inputs.PaginationInput true "Post id"
// @Success 200 {array} models.Post
// @Router /posts [get]
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

// Find post by id godoc
// @Summary Find post by id
// @Schemes
// @Description Find one post by id
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post id"
// @Success 200 {object} models.Post
// @Router /posts/{id} [get]
func FindOnePost(context *gin.Context) {
	var post models.Post
	postId := context.Param("id")

	err := cache.GetFromCache(POST_CACHE_PREFIX+postId, post)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"data": post})
		return
	}

	if err := models.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
	}

	cache.SaveToCache(POST_CACHE_PREFIX+postId, post, time.Hour)

	context.JSON(http.StatusOK, gin.H{"data": post})
}
