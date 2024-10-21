package db

import (
	"database/sql"
	"kaiquecaires/real-time-leaderboard/cmd/models"
)

type GameStore interface {
	Insert(models.CreateGameParams) (*models.Game, error)
	Get() ([]models.Game, error)
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

func (s *PostgresGameStore) Get() ([]models.Game, error) {
	query := "SELECT * FROM games"
	var games []models.Game
	rows, err := s.conn.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var game models.Game
		err := rows.Scan(&game.Id, &game.Name)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}
