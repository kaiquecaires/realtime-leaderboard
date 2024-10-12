package main

import (
	"kaiquecaires/real-time-leaderboard/cmd/databases"
	"kaiquecaires/real-time-leaderboard/cmd/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "API IS ON FIRE!")
	})

	conn := databases.GetPostgresInstance()
	userStore := databases.NewPostgresUserStore(conn)
	signUpHandler := handlers.NewSignUpHandler(userStore)
	route.POST("/signup", signUpHandler.Handle)

	gameStore := databases.NewPostgresGameStore(conn)
	createGameHandler := handlers.NewGameHandler(gameStore)
	route.POST("/game", createGameHandler.CreateGameHandler)

	route.Run("0.0.0.0:8080")
}
