package handlers

import (
	"kaiquecaires/real-time-leaderboard/cmd/db"
	"kaiquecaires/real-time-leaderboard/cmd/messaging"
	"kaiquecaires/real-time-leaderboard/cmd/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserScoreHandler struct {
	userScorePublisher messaging.UserScorePublisher
	userScoreStore     db.UserScoreStore
	leaderboardCache   db.LeaderboardCache
}

func NewUserScoreHandler(userScorePublisher messaging.UserScorePublisher, userScoreStore db.UserScoreStore, leaderboardCache db.LeaderboardCache) *UserScoreHandler {
	return &UserScoreHandler{userScorePublisher: userScorePublisher, userScoreStore: userScoreStore, leaderboardCache: leaderboardCache}
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
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"message": "score successfully sent"})
}

func (h *UserScoreHandler) HandleGetLeaderboard(c *gin.Context) {
	var getLeaderboardParams models.GetLeaderboardParams

	if err := c.ShouldBindQuery(&getLeaderboardParams); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	cachedLeaderboard, err := h.leaderboardCache.Get(c.Request.Context())

	if err != nil {
		log.Println("Failed to get leaderboard  from cache!", err)
	}

	if cachedLeaderboard != nil && len(cachedLeaderboard) > 0 {
		log.Println("got from cache!")
		c.JSON(http.StatusOK, cachedLeaderboard)
		return
	}

	leaderboard, err := h.userScoreStore.GetLeaderboard(getLeaderboardParams)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leaderboard)
}
