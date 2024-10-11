package handlers

import (
	"fmt"
	"kaiquecaires/real-time-leaderboard/cmd/databases"
	"kaiquecaires/real-time-leaderboard/cmd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpHandler struct {
	userStore databases.UserStore
}

func NewSignUpHandler(userStore databases.UserStore) *SignUpHandler {
	return &SignUpHandler{
		userStore: userStore,
	}
}

func (h *SignUpHandler) Handle(c *gin.Context) {
	var createUserParams models.CreateUserParams

	if err := c.ShouldBind(&createUserParams); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Error on signup: %v", err)})
		return
	}

	if errors := createUserParams.Validate(); len(errors) > 0 {
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	user, err := h.userStore.InsertUser(createUserParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
