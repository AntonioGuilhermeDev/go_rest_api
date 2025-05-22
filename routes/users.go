package routes

import (
	"net/http"

	"github.com/AntonioGuilhermeDev/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": "Could not parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": "Could not create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User create successfully!"})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}
