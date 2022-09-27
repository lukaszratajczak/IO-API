package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizAPP/pkg/config"
	"quizAPP/pkg/models"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := config.GetDB().Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	var score models.Score
	score.User = user
	score.Score = 0
	newscore := config.GetDB().Create(&score)
	if newscore.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": newscore.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})

}
