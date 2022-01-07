package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mustafakocatepe/find-gaming-friends/model"
	userRepository "github.com/mustafakocatepe/find-gaming-friends/repository/user"
	"github.com/mustafakocatepe/find-gaming-friends/utils"
	"golang.org/x/crypto/bcrypt"
)

var users []model.UserDTO

type Controller struct{}

func (c Controller) GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.UserDTO
		var error model.Error

		users = []model.UserDTO{}
		userRepo := userRepository.UserRepository{}

		users, err := userRepo.GetUsers(db, user, users)

		if err != nil {
			error.Message = "Server Error"
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.RespondWithJSON(w, users, http.StatusOK)
	}
}

func (c Controller) AddUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createUserDTO model.CreateUserDTO
		var error model.Error

		json.NewDecoder(r.Body).Decode(&createUserDTO)

		err := createUserDTO.Validate()

		if err != nil {
			error.Message = err.Error()
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		user := model.User{
			UserName:  createUserDTO.Username,
			Email:     createUserDTO.Email,
			Password:  createUserDTO.Password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		user.Password = string(hashed)

		userRepo := userRepository.UserRepository{}
		userId, err := userRepo.AddUser(db, user)

		if err != nil {
			error.Message = "Server Error"
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.RespondWithJSON(w, userId, http.StatusOK)
	}
}

func (c Controller) RemoveUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error model.Error
		params := mux.Vars(r)

		userRepo := userRepository.UserRepository{}

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			error.Message = "Incorrect id."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		rowsAffected, err := userRepo.RemoveUser(db, id)

		if err != nil {
			error.Message = "Server error"
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.RespondWithJSON(w, rowsAffected, http.StatusNoContent)
	}
}

func (c Controller) UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userDTO model.UpdateUserDTO
		var error model.Error

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			error.Message = "Incorrect id."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		json.NewDecoder(r.Body).Decode(&userDTO)

		var user model.User
		user.Id = id

		if len(userDTO.UserName) != 0 {
			user.UserName = userDTO.UserName
		}

		if len(userDTO.Email) != 0 {
			user.Email = userDTO.Email
		}

		if len(userDTO.Bio) != 0 {
			user.Bio = userDTO.Bio
		}

		userRepo := userRepository.UserRepository{}
		_, err = userRepo.UpdateUser(db, user)

		if err != nil {
			error.Message = "Server error"
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")

		utils.RespondWithJSON(w, "", http.StatusNoContent)
	}
}
