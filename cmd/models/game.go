package models

import "errors"

type Game struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateGameParams struct {
	Name string `json:"name"`
}

func (p *CreateGameParams) Validate() error {
	if len(p.Name) < 5 {
		return errors.New("name must have at least 5 characters")
	}

	if len(p.Name) > 50 {
		return errors.New("name must have at most 50 characters")
	}

	return nil
}
