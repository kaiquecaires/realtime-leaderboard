package databases

import (
	"database/sql"
	"kaiquecaires/real-time-leaderboard/cmd/models"

	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	InsertUser(models.CreateUserParams) (*models.User, error)
	GetByUsername(string) (*models.User, error)
}

type PostgresUserStore struct {
	conn *sql.DB
}

func NewPostgresUserStore(conn *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		conn: conn,
	}
}

func (s *PostgresUserStore) InsertUser(params models.CreateUserParams) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	var userId int
	err = s.conn.QueryRow(query, params.Username, hashedPassword).Scan(&userId)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:       userId,
		Username: params.Username,
		Password: params.Password,
	}, err
}

func (s *PostgresUserStore) GetByUsername(username string) (*models.User, error) {
	var id int
	var password string

	query := "SELECT id, password FROM users WHERE username = $1"
	err := s.conn.QueryRow(query, username).Scan(&id, &password)
	return &models.User{
		Id:       id,
		Username: username,
		Password: password,
	}, err
}
