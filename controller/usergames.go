package controller

import (
	"database/sql"
	"encoding/json"
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

		userGamesRepo := userGamesRepository.UserGamesRepository{}

		userId, err := strconv.Atoi(params["id"])

		if err != nil {
			error.Message = "Incorrect id."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		userGames := model.UserGamesDTO{}

		userGames, err = userGamesRepo.GetUserGames(db, userId, userGames)

		if err != nil {
			error.Message = "Server error"
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.RespondWithJSON(w, userGames, http.StatusOK)
	}
}

func (c Controller) AddUserGames(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var createUserGamesDTO model.CreateUserGamesDTO
		var error model.Error

		json.NewDecoder(r.Body).Decode(&createUserGamesDTO)

		userGamesRepo := userGamesRepository.UserGamesRepository{}
		err := userGamesRepo.AddUserGames(db, createUserGamesDTO.UserId, createUserGamesDTO.GameId)

		if err != nil {
			error.Message = "Server Error"
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.RespondWithJSON(w, "", http.StatusOK)
	}
}
