package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizAPP/pkg/auth"
	"quizAPP/pkg/config"
	"quizAPP/pkg/models"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})

}

func GetScore(context *gin.Context) {
	NewScore := models.GetEveryScore()
	context.JSON(http.StatusOK, gin.H{"data": NewScore})
}

func CreateScore(context *gin.Context) {
	var CreateScore models.Score

	tokenString := context.GetHeader("Authorization")
	userEmail := auth.Getuserfromtoken(tokenString)
	myUser := models.User{}
	db := config.GetDB().Where("Email=?", userEmail).Find(&myUser)
	if db.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": db.Error.Error()})
		context.Abort()
		return
	}
	CreateScore.User = myUser
	if err := context.ShouldBindJSON(&CreateScore); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := config.GetDB().Create(&CreateScore)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": &CreateScore.User.Username, "score": &CreateScore.Score})
}
func GetScoreByUser(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	userEmail := auth.Getuserfromtoken(tokenString)

	UserScore, _ := models.GetScoreByUser(userEmail)
	context.JSON(http.StatusOK, gin.H{"score": &UserScore.Score, "user": &UserScore.User.Username})
}
func UpdateUserScore(context *gin.Context) {
	var updateScore models.Score
	if err := context.ShouldBindJSON(&updateScore); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	tokenString := context.GetHeader("Authorization")
	userEmail := auth.Getuserfromtoken(tokenString)
	scoreDetails, _ := models.GetScoreByUser(userEmail)
	test := config.GetDB().Debug().Model(&scoreDetails).Where("user_id = ?", scoreDetails.UserID).Update("Score", scoreDetails.Score+updateScore.Score)
	if test.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": test.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": scoreDetails.User.Username, "score": scoreDetails.Score})
}
