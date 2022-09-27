package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"quizAPP/pkg/config"
	"quizAPP/pkg/middlewares"
	"quizAPP/pkg/models"
	"strconv"
)

func GetQuestionN(context *gin.Context) {
	newQuestion := models.GetAllQuestions()
	context.JSON(http.StatusOK, gin.H{"data": newQuestion})
}
func GetSomeQuestionsN(context *gin.Context) {
	quantity := context.Param("quantity")
	q, err := strconv.ParseInt(quantity, 0, 0)
	if err != nil {
		fmt.Println("error while parsing quantity")
	}
	subject := context.Param("subject")
	firstyear := context.Param("first-year")
	f, err := strconv.ParseInt(firstyear, 0, 0)
	if err != nil {
		fmt.Println("error while parsing first year")
	}
	lastyear := context.Param("last-year")
	l, err := strconv.ParseInt(lastyear, 0, 0)
	if err != nil {
		fmt.Println("error while parsing last year")
	}
	Questions := models.GetSomeQuestions(int(q), subject, f, l)
	context.JSON(http.StatusOK, gin.H{"data": Questions})

}
func GetRankedQuestionsN(context *gin.Context) {
	Questions := models.GetSomeQuestions(10, "Wszystkie", -1, 9999999)
	context.JSON(http.StatusOK, gin.H{"data": Questions})

}
func GetRandomQuestionN(context *gin.Context) {
	randomQuestion, _ := models.GetRandomQuestion()
	context.JSON(http.StatusOK, gin.H{"data": randomQuestion})
}
func GetQuestionByIdN(context *gin.Context) {
	ID := context.Param("questionId")
	id, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing ")
	}
	QuestionDetails, _ := models.GetQuestionById(id)
	context.JSON(http.StatusOK, gin.H{"data": QuestionDetails})
}
func CreateQuestionN(context *gin.Context) {
	var CreateQuestion models.Question
	if err := context.ShouldBindJSON(&CreateQuestion); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := config.GetDB().Create(&CreateQuestion)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": CreateQuestion})
}

func SaveImage(context *gin.Context) {
	context.Request.ParseMultipartForm(5 << 10)
	file, handler, err := context.Request.FormFile("myFile")
	if err != nil {
		fmt.Println("error retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded Failed: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	os.MkdirAll("question-images", os.ModePerm)
	tempFile, err := ioutil.TempFile("question-images", "question-*.PNG")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	fmt.Println(tempFile.Name())

	middlewares.UploadFile(tempFile.Name())
	tempFile.Close()
	err = os.RemoveAll(tempFile.Name())
	if err != nil {
		fmt.Println(err)
	}
	filenametodb := "https://lratajczakmybucketforquestions.s3.amazonaws.com/" + tempFile.Name()
	fmt.Println(filenametodb)
	context.Data(http.StatusOK, "string", []byte(filenametodb))
}

func DeleteQuestion(context *gin.Context) {
	ID := context.Param("questionId")
	id, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	models.DeleteQuestion(id)
	context.JSON(http.StatusOK, gin.H{"deleted": ID})
}

func UpdateQuestion(context *gin.Context) {
	var updateQuestion models.Question
	if err := context.ShouldBindJSON(&updateQuestion); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	ID := context.Param("questionId")
	id, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	questionDetails, db := models.GetQuestionById(id)
	if updateQuestion.Question != "" {
		questionDetails.Question = updateQuestion.Question
	}
	if updateQuestion.AnswerA != "" {
		questionDetails.AnswerA = updateQuestion.AnswerA
	}
	if updateQuestion.AnswerB != "" {
		questionDetails.AnswerB = updateQuestion.AnswerB
	}
	if updateQuestion.AnswerC != "" {
		questionDetails.AnswerC = updateQuestion.AnswerC
	}
	if updateQuestion.AnswerD != "" {
		questionDetails.AnswerD = updateQuestion.AnswerD
	}
	if updateQuestion.CorrectAnswer != "" {
		questionDetails.CorrectAnswer = updateQuestion.CorrectAnswer
	}
	if updateQuestion.Subject != "" {
		questionDetails.Subject = updateQuestion.Subject
	}
	if updateQuestion.Year != 0 {
		questionDetails.Year = updateQuestion.Year
	}
	db.Save(&questionDetails)
	context.JSON(http.StatusOK, gin.H{"data": questionDetails})
}
