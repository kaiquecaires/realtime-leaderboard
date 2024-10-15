package handlers

import (
	"kaiquecaires/real-time-leaderboard/cmd/auth"
	"kaiquecaires/real-time-leaderboard/cmd/db"
	"kaiquecaires/real-time-leaderboard/cmd/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	userStore db.UserStore
}

func NewLoginHandler(userStore db.UserStore) *LoginHandler {
	return &LoginHandler{userStore: userStore}
}

func (h *LoginHandler) Handle(c *gin.Context) {
	var loginParams models.LoginParams

	if err := c.ShouldBind(&loginParams); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if errors := loginParams.Validate(); len(errors) > 0 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errors})
		return
	}

	user, err := h.userStore.GetByUsername(loginParams.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginParams.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid password"})
		return
	}

	token, err := auth.GenerateToken(user.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Login{AccessToken: token})
}
