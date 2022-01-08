package gameRepository

import (
	"database/sql"

	"github.com/mustafakocatepe/find-gaming-friends/model"
)

type GameRepository struct{}

func (u GameRepository) GetGames(db *sql.DB, gameDTO model.GameDTO, gameDTOs []model.GameDTO) ([]model.GameDTO, error) {
	rows, err := db.Query("SELECT id, name, developer, publisher FROM games WHERE is_active = true ")
	if err != nil {
		return []model.GameDTO{}, err
	}

	for rows.Next() {
		err = rows.Scan(&gameDTO.Id, &gameDTO.Name, &gameDTO.Developer, &gameDTO.Publisher)
		gameDTOs = append(gameDTOs, gameDTO)
	}

	if err != nil {
		return []model.GameDTO{}, err
	}
	return gameDTOs, nil
}

func (u GameRepository) AddGame(db *sql.DB, game model.Game) (int, error) {
	query := "INSERT INTO games (name, developer, publisher, created_at, updated_at, is_active) VALUES($1,$2,$3,$4,$5,$6) RETURNING id"
	err := db.QueryRow(query, game.Name, game.Developer, game.Publisher, game.CreatedAt, game.UpdatedAt, true).Scan(&game.Id)

	if err != nil {
		return 0, nil
	}

	return int(game.Id), nil
}
