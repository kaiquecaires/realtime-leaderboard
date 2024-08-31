package main

import (
	"kaiquecaires/real-time-leaderboard/cmd/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "API IS ON FIRE!")
	})
	route.POST("/signup", handlers.SignUp)
	route.Run("0.0.0.0:8080")
}
