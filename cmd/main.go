package main

import (
	"kaiquecaires/real-time-leaderboard/cmd/db"
	"kaiquecaires/real-time-leaderboard/cmd/handlers"
	"kaiquecaires/real-time-leaderboard/cmd/messaging"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "API IS ON FIRE!")
	})

	conn := db.GetPostgresInstance()
	userStore := db.NewPostgresUserStore(conn)
	signUpHandler := handlers.NewSignUpHandler(userStore)
	route.POST("/signup", signUpHandler.Handle)

	gameStore := db.NewPostgresGameStore(conn)
	createGameHandler := handlers.NewGameHandler(gameStore)
	route.POST("/game", createGameHandler.CreateGameHandler)

	producer := messaging.GetProducer()
	userScorePublisher := messaging.NewKafkaUserScorePublisher(producer)
	userScoreHandler := handlers.NewUserScoreHandler(userScorePublisher)
	route.POST("/user-score", userScoreHandler.HandleSendUserScore)

	userScoreStore := db.NewPostgresUserScoreStore(conn)
	leaderboardConsumer := messaging.NewLeaderboardConsumer(userScoreStore)
	go leaderboardConsumer.Consume("leaderboard_postgres_1", "leaderdoard_postgres")

	loginHandler := handlers.NewLoginHandler(userStore)
	route.POST("/login", loginHandler.Handle)

	route.Run("0.0.0.0:8080")
}
