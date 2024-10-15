package db

import (
	"database/sql"
	"kaiquecaires/real-time-leaderboard/cmd/models"
)

type UserScoreStore interface {
	Insert(models.CreateUserScoreParams) error
}

type PostgresUserScoreStore struct {
	conn *sql.DB
}

func NewPostgresUserScoreStore(conn *sql.DB) *PostgresUserScoreStore {
	return &PostgresUserScoreStore{conn: conn}
}

func (s *PostgresUserScoreStore) Insert(params models.CreateUserScoreParams) error {
	query := "INSERT INTO user_scores (user_id, game_id, score) VALUES ($1, $2, $3) RETURNING id"
	var userScoreId int
	err := s.conn.QueryRow(query, params.UserId, params.GameId, params.Score).Scan(&userScoreId)
	return err
}
