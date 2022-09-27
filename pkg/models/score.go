package models

import (
	"github.com/jinzhu/gorm"
	"quizAPP/pkg/config"
)

type Score struct {
	UserID int
	User   User
	Score  int `json:"score" gorm:"default:'0'"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Score{})
}

func GetEveryScore() []Score {
	var Score []Score
	//db.Find(&Score)
	db.Preload("User").Find(&Score)
	return Score
}

func GetScoreByUser(email string) (*Score, *gorm.DB) {
	var getScore Score
	//db.Debug().Joins("User", db.Where(&User{Email: email})).Find(&getScore)
	//db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
	//db.Model(&Score{}).Joins("inner join users on users.ID = scores.user_id").Scan(&getScore)
	db.Debug().Preload("User").Joins("Join users on users.ID = scores.user_id AND users.Email = ?", email).Find(&getScore)
	return &getScore, db
}
