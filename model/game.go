package model

import "time"

type Game struct {
	GameId    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Developer string    `json:"developer"`
	Publisher string    `json:"publisher"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
