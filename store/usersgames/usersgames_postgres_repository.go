package usersgames

import (
	"database/sql"
	"time"

	"github.com/mustafakocatepe/find-gaming-friends/model"
)

type UserGamesRepository struct{}

func (u UserGamesRepository) GetUserGames(db *sql.DB, id int, userGames model.UserGamesDTO) (model.UserGamesDTO, error) {
	var gameDTO model.GameDTO

	rows, err := db.Query("SELECT game_id FROM UserGames WHERE is_active = true and user_id =$1", id)
	if err != nil {
		return model.UserGamesDTO{}, err
	}

	usersgames := model.UserGames{}

	for rows.Next() {
		err = rows.Scan(&usersgames.GameId)

		gameRows, _ := db.Query("SELECT id, name, developer, publisher FROM Games WHERE is_active = true and id =$1", &usersgames.GameId)

		for gameRows.Next() {
			err = gameRows.Scan(&gameDTO.Id, &gameDTO.Name, &gameDTO.Developer, &gameDTO.Publisher)
		}
		userGames.GameList = append(userGames.GameList, gameDTO)
	}

	userGames.UserId = id

	if err != nil {
		return model.UserGamesDTO{}, err
	}
	return userGames, nil
}

func (u UserGamesRepository) AddUserGames(db *sql.DB, userId int, gameId int) error {

	query := "INSERT INTO usergames (user_id, game_id, created_at, updated_at, is_active) VALUES($1,$2,$3,$4,$5)"

	err := db.QueryRow(query, userId, gameId, time.Now(), time.Now(), true)

	if err != nil {
		return nil
	}

	return nil
}
