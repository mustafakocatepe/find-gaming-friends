package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mustafakocatepe/find-gaming-friends/model"
	userGamesRepository "github.com/mustafakocatepe/find-gaming-friends/repository/usersgames"
	"github.com/mustafakocatepe/find-gaming-friends/utils"
)

func (c Controller) GetUserGames(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error model.Error
		params := mux.Vars(r)

		userGamesRepo := userGamesRepository.UsersGamesRepository{}

		userId, err := strconv.Atoi(params["id"])

		if err != nil {
			error.Message = "Incorrect id."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		userGames := model.UserGamesDTO{}

		userGames, err = userGamesRepo.GetUserGames(db, userId, userGames)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		utils.JSON(w, userGames, http.StatusOK)
	}
}
