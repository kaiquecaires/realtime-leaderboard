package databases

import (
	"database/sql"
	"kaiquecaires/real-time-leaderboard/cmd/models"
)

type GameStore interface {
	Insert(models.CreateGameParams) (*models.Game, error)
}

type PostgresGameStore struct {
	conn *sql.DB
}

func NewPostgresGameStore(conn *sql.DB) *PostgresGameStore {
	return &PostgresGameStore{conn: conn}
}

func (s *PostgresGameStore) Insert(params models.CreateGameParams) (*models.Game, error) {
	query := "INSERT INTO games (name) VALUES ($1) RETURNING id"
	var gameId int
	err := s.conn.QueryRow(query, params.Name).Scan(&gameId)

	if err != nil {
		return nil, err
	}

	return &models.Game{
		Id:   gameId,
		Name: params.Name,
	}, nil
}
