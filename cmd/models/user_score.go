package models

import "errors"

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

	return nil
}

type UserScore struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	GameId int `json:"game_id"`
	Score  int `json:"score"`
}
