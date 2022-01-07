package model

import (
	"time"
)

type UserGames struct {
	Id        uint      `json:"id"`
	UserId    string    `json:"user_id"`
	GameId    string    `json:"game_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

type UserGamesDTO struct {
	UserId   int       `json:"user_id"`
	GameList []GameDTO `json:"game_list"`
}

type CreateUserGamesDTO struct {
	UserId int `json:"user_id"`
	GameId int `json:"game_id"`
}
