package controllers

import (
	"net/http"

	"github.com/StasJDM/go-gin-gorm-api/helpers"
	"github.com/StasJDM/go-gin-gorm-api/models"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input models.RegisterInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
	var input models.LoginInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helpers.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
