package usersgames

import (
	"database/sql"

	"github.com/mustafakocatepe/find-gaming-friends/model"
)

type UsersGamesRepository struct{}

func (u UsersGamesRepository) GetUserGames(db *sql.DB, id int, userGames model.UserGamesDTO) (model.UserGamesDTO, error) {
	var gameDTO model.GameDTO

	rows, err := db.Query("SELECT game_id FROM UsersGames WHERE is_active = true and user_id =$1", id)
	if err != nil {
		return model.UserGamesDTO{}, err
	}

	usersgames := model.UsersGames{}

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
