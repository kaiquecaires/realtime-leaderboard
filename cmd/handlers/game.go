package handlers

import (
	"kaiquecaires/real-time-leaderboard/cmd/db"
	"kaiquecaires/real-time-leaderboard/cmd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gameStore db.GameStore
}

func NewGameHandler(gameStore db.GameStore) *GameHandler {
	return &GameHandler{gameStore: gameStore}
}

func (h *GameHandler) CreateGameHandler(c *gin.Context) {
	var createGameParams models.CreateGameParams

	if err := c.ShouldBind(&createGameParams); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := createGameParams.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	game, err := h.gameStore.Insert(createGameParams)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}
