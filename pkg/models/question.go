package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
	"quizAPP/pkg/config"
	"time"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var db *gorm.DB

type Question struct {
	gorm.Model
	Question      string `gorm:""json:"question"`
	AnswerA       string `gorm:""json:"answerA"`
	AnswerB       string `gorm:""json:"answerB"`
	AnswerC       string `gorm:""json:"answerC"`
	AnswerD       string `gorm:""json:"answerD"`
	CorrectAnswer string `gorm:""json:"correctAnswer"`
	Subject       string `gorm:""json:"subject"`
	Year          int64  `gorm:""json:"year"`
	ImageLink     string `gorm:""json:"ImageLink"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Question{})
}

func (q *Question) CreateQuestion() *Question {
	db.NewRecord(q)
	db.Create(&q)
	return q
}

func GetAllQuestions() []Question {
	var Question []Question
	db.Find(&Question)
	return Question
}

func GetQuestionById(Id int64) (*Question, *gorm.DB) {
	var getQuestion Question
	db := db.Where("ID=?", Id).Find(&getQuestion)
	return &getQuestion, db
}

func GetSomeQuestions(quantity int, subject string, firstyear int64, lastyear int64) []Question {
	var randindexes []int
	var ids []int64
	var Question []Question
	if subject != "Wszystkie" {
		db.Where("subject = ? AND year BETWEEN ? and ? ", subject, firstyear, lastyear).Find(&Question).Pluck("ID", &ids)
	} else {
		db.Where("year BETWEEN ? and ? ", firstyear, lastyear).Find(&Question).Pluck("ID", &ids)
	}
	rand.Seed(time.Now().Unix())
	for len(randindexes) < quantity {
		var randomid = int(ids[rand.Intn(len(ids))])

		if !contains(randindexes, randomid) {
			randindexes = append(randindexes, randomid)

		}

	}
	fmt.Println(randindexes)
	db.Find(&Question, randindexes)
	return Question
}

func GetRandomQuestion() (*Question, *gorm.DB) {
	var LastQuestion Question
	db.Last(&LastQuestion)
	var getQuestion Question

	exist := true
	for exist {
		randomid := int(rand.Intn(int(LastQuestion.ID-1) + 1))
		db := db.Where("ID=?", randomid).Find(&getQuestion)
		if getQuestion.Question != "" {
			return &getQuestion, db
		}
	}
	return &getQuestion, db
}

func DeleteQuestion(ID int64) Question {
	var Question Question
	db.Where("ID=?", ID).Delete(Question)
	return Question
}
