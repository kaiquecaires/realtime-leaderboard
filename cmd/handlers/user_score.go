package handlers

import (
	"kaiquecaires/real-time-leaderboard/cmd/messaging"
	"kaiquecaires/real-time-leaderboard/cmd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserScoreHandler struct {
	userScorePublisher messaging.UserScorePublisher
}

func NewUserScoreHandler(userScorePublisher messaging.UserScorePublisher) *UserScoreHandler {
	return &UserScoreHandler{userScorePublisher: userScorePublisher}
}

func (h *UserScoreHandler) HandleSendUserScore(c *gin.Context) {
	var createUserScore models.CreateUserScoreParams

	if err := c.ShouldBind(&createUserScore); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := createUserScore.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err := h.userScorePublisher.NewScore(createUserScore)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, map[string]string{"message": "score successfully sent"})
}
