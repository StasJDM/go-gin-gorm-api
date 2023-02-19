package controllers

import (
	"net/http"

	"github.com/StasJDM/go-gin-gorm-api/helpers"
	"github.com/StasJDM/go-gin-gorm-api/inputs"
	"github.com/StasJDM/go-gin-gorm-api/models"
	"github.com/gin-gonic/gin"
)

type JWTResponse struct {
	JWT string `json:"jwt"`
}

// Register godoc
// @Summary Register
// @Schemes
// @Description Register user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body inputs.RegisterInput true "Register data"
// @Success 200 {object} JWTResponse "jwt"
// @Router /auth/register [post]
func Register(context *gin.Context) {
	var input inputs.RegisterInput

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

// Login godoc
// @Summary Login
// @Schemes
// @Description Login using username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body inputs.LoginInput true "Login data"
// @Success 200 {object} JWTResponse "jwt"
// @Router /auth/login [post]
func Login(context *gin.Context) {
	var input inputs.LoginInput

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
