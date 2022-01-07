package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/mustafakocatepe/find-gaming-friends/model"
	gameRepository "github.com/mustafakocatepe/find-gaming-friends/repository/game"
	"github.com/mustafakocatepe/find-gaming-friends/utils"
)

var games []model.GameDTO

func (c Controller) GetGames(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var game model.GameDTO
		var error model.Error

		games = []model.GameDTO{}
		gameRepo := gameRepository.GameRepository{}

		games, err := gameRepo.GetGames(db, game, games)

		if err != nil {
			error.Message = "Server Error"
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.RespondWithJSON(w, games, http.StatusOK)
	}
}

func (c Controller) AddGame(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createGameDTO model.CreateGameDTO
		var error model.Error

		json.NewDecoder(r.Body).Decode(&createGameDTO)

		err := createGameDTO.Validate()

		if err != nil {
			error.Message = err.Error()
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		game := model.Game{
			Name:      createGameDTO.Name,
			Developer: createGameDTO.Developer,
			Publisher: createGameDTO.Publisher,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		gameRepo := gameRepository.GameRepository{}
		gameId, err := gameRepo.AddGame(db, game)

		if err != nil {
			error.Message = "Server Error"
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.RespondWithJSON(w, gameId, http.StatusOK)
	}
}
