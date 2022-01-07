package model

import (
	"errors"
	"time"
)

// Game is our main model for Games
type Game struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Developer string    `json:"developer"`
	Publisher string    `json:"publisher"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

// GameDTO is our data transfer object for Game
type GameDTO struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Developer string `json:"developer"`
	Publisher string `json:"publisher"`
}

type CreateGameDTO struct {
	Name      string `json:"name"`
	Developer string `json:"developer"`
	Publisher string `json:"publisher"`
}

func (game *CreateGameDTO) Validate() error {

	if len(game.Name) == 0 {
		return errors.New("name should not be empty")
	}

	if len(game.Developer) == 0 {
		return errors.New("developer should not be empty")
	}

	if len(game.Publisher) == 0 {
		return errors.New("publisher should not be empty")
	}

	return nil
}
