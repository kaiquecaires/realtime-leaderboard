package db

import (
	"database/sql"
	"fmt"
	"kaiquecaires/real-time-leaderboard/cmd/models"
)

type UserScoreStore interface {
	Insert(models.CreateUserScoreParams) error
	GetLeaderboard(models.GetLeaderboardParams) ([]models.Leaderboard, error)
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

func (s *PostgresUserScoreStore) GetLeaderboard(params models.GetLeaderboardParams) ([]models.Leaderboard, error) {
	query := `
  SELECT username, SUM(score) as total_score FROM user_scores us
  INNER JOIN users u ON u.id = us.user_id
  GROUP BY u.id ORDER BY total_score DESC
  LIMIT $1 OFFSET $2;

  `

	if params.Limit == 0 {
		params.Limit = 10
	}

	rows, err := s.conn.Query(query, params.Limit, params.Offset)

	if err != nil {
		return nil, err
	}

	var leaderboard []models.Leaderboard

	for rows.Next() {
		var user models.Leaderboard
		err := rows.Scan(&user.Username, &user.Score)
		if err != nil {
			return nil, err
		}
		leaderboard = append(leaderboard, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return leaderboard, nil
}
