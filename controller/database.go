package controller

import (
	"database/sql"
	"net/http"

	"github.com/mustafakocatepe/find-gaming-friends/db/migrate/postgres"
	"github.com/mustafakocatepe/find-gaming-friends/model"
	"github.com/mustafakocatepe/find-gaming-friends/utils"
)

func (c Controller) CreateDatabaseTables(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var error model.Error
		err := postgres.Migrate(db)

		if err != nil {
			error.Message = "Database could not be created"
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.RespondWithJSON(w, "Database Created", http.StatusOK)
	}
}
