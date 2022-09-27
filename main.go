package main

import (
	"IO-API/pkg/controllers"
	"IO-API/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	router := initRouter()
	router.Run("0.0.0.0:8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{

		api.GET("/questions/", controllers.GetQuestionN)
		//Dostajemy wszystkie pytania które są w baize
		api.GET("/question/:questionId", controllers.GetQuestionByIdN)
		//dostajemy konkretne pytanie, odwołujemy sie po id pytania
		api.POST("/questions", controllers.CreateQuestionN)
		//dodajemy pytanie
		api.POST("/addimage", controllers.SaveImage)
		api.DELETE("/questions/:questionId", controllers.DeleteQuestion)
		//usuwamy pytanie
		api.PUT("/questions/:questionId", controllers.UpdateQuestion)
		//zmieniamy pytanie
		api.GET("/questions/:quantity/:subject/:first-year/:last-year", controllers.GetSomeQuestionsN)
		//dostajemy konkretną ilość pytań z konkretnego przedmiotu z konkretnego przedzialu lat
		api.GET("/questions/ranked", controllers.GetRankedQuestionsN)
		api.GET("/questions/random", controllers.GetRandomQuestionN)
		//randomowo dostajemy jedno pytanie z calej puli

		api.GET("/scores/", controllers.GetScore)
		//dostajemy wyniki wszystkich użytkowików
		api.POST("/token", controllers.GenerateToken)
		//logowanie do bazy -> dostajemy token który definiuje uzytkownika przez 60 minut
		api.POST("/user/register", controllers.RegisterUser)
		//rejstracja uzytkownika
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			//test czy dziala token
			secured.POST("/scores", controllers.CreateScore)
			//tworzymy wynik dla uzytkownika - chyba po chuju ogolnie bo to sie tworzy przy rejestracji
			secured.GET("/score", controllers.GetScoreByUser)
			//dostajemy wynik dla konkretnego uzytkokwnika ktorego definiujemy przez token
			secured.PATCH("/score", controllers.UpdateUserScore)
			//updatujemy wynik dla konkretnego uzytkownika ktorego definiujemy przez token
		}
	}
	return router
}
