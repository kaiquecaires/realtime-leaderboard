package models

import (
	"errors"
)

type CreateUserScoreParams struct {
	UserId int `json:"user_id"`
	GameId int `json:"game_id"`
	Score  int `json:"score"`
}

func (p *CreateUserScoreParams) Validate() error {
	if p.Score < 0 {
		return errors.New("score cannot be smaller than 0")
	}

	if p.Score > 100 {
		return errors.New("score cannot be grater than 100")
	}

	if p.UserId == 0 {
		return errors.New("user_id is required")
	}

	if p.GameId == 0 {
		return errors.New("game_id is required")
	}

	return nil
}

type UserScore struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	GameId int `json:"game_id"`
	Score  int `json:"score"`
}

type GetLeaderboardParams struct {
	Limit  int
	Offset int
}

type Leaderboard struct {
	Username string `json:"user_id:"`
	Score    int    `json:"score"`
}
