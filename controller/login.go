package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/mustafakocatepe/find-gaming-friends/model"
	userRepository "github.com/mustafakocatepe/find-gaming-friends/repository/user"
	"github.com/mustafakocatepe/find-gaming-friends/utils"
)

func (c Controller) LoginControl(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginDTO model.LoginDTO
		var error model.Error

		json.NewDecoder(r.Body).Decode(&loginDTO)

		err := loginDTO.Validate()

		if err != nil {
			error.Message = err.Error()
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		/*hashed, _ := bcrypt.GenerateFromPassword([]byte(loginDTO.Password), 8)
		loginDTO.Password = string(hashed)*/

		bookRepo := userRepository.UserRepository{}
		success, err := bookRepo.Login(db, loginDTO)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if success {
			utils.JSON(w, "Login successful ", http.StatusOK)
		} else {
			error.Message = "Login Failed"
			utils.SendError(w, http.StatusUnauthorized, error)
		}
	}
}
