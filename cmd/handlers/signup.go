package handlers

import (
	"fmt"
	"kaiquecaires/real-time-leaderboard/cmd/databases"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignUpBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context) {
	var user SignUpBody

	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error on signup: %v", err))
		return
	}

	if user.Username == "" || user.Password == "" {
		c.String(http.StatusBadRequest, "username and password must be provided")
		return
	}

	if len(user.Username) < 5 || len(user.Username) > 20 {
		c.String(http.StatusBadRequest, "username must have between 5 and 20 chars")
		return
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		c.String(http.StatusBadRequest, "password must have between 8 and 20 chars")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error on signup: %v", err))
		return
	}

	conn := databases.GetDBInstance()

	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = conn.Exec(query, user.Username, hashedPassword)

	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error inserting user: %v", err))
		return
	}

	c.String(200, "User created!")
}
