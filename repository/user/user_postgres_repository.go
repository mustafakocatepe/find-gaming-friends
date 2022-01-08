package userRepository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mustafakocatepe/find-gaming-friends/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func (u UserRepository) GetUsers(db *sql.DB, userDTO model.UserDTO, userDTOs []model.UserDTO) ([]model.UserDTO, error) {
	rows, err := db.Query("SELECT id, username, email, bio FROM Users WHERE is_active = true ")
	if err != nil {
		return []model.UserDTO{}, err
	}

	for rows.Next() {
		err = rows.Scan(&userDTO.Id, &userDTO.UserName, &userDTO.Email, &userDTO.Bio)
		userDTOs = append(userDTOs, userDTO)
	}

	if err != nil {
		return []model.UserDTO{}, err
	}
	return userDTOs, nil
}

func (u UserRepository) AddUser(db *sql.DB, user model.User) (int, error) {
	query := "INSERT INTO users (email, password, username, bio, image, created_at, updated_at, is_active) VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id"
	err := db.QueryRow(query, user.Email, user.Password, user.UserName, user.Bio, user.Image, user.CreatedAt, user.UpdatedAt, true).Scan(&user.Id)

	if err != nil {
		return 0, nil
	}

	return int(user.Id), nil
}

func (U UserRepository) RemoveUser(db *sql.DB, id int) (int64, error) {

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec("UPDATE Users SET is_active = false WHERE id = $1", id)

	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, err
	}
	var count int

	resultCount, _ := tx.Query("SELECT COUNT(*) FROM UserGames WHERE user_id = $1 AND is_active = true", id)

	for resultCount.Next() {
		_ = resultCount.Scan(count)
	}

	if count > 0 {
		_, err2 := tx.Exec("UPDATE UserGames SET is_active = false WHERE user_id = $1", id)

		if err2 != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				return 0, rollbackErr
			}
			return 0, err2
		}

	}

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (u UserRepository) Login(db *sql.DB, loginDTO model.LoginDTO) (bool, error) {

	var user model.User
	query := "SELECT email, password FROM Users where email = $1 and is_active = true"
	err := db.QueryRow(query, loginDTO.Email).Scan(&user.Email, &user.Password)

	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))

	if err == nil {
		return true, err
	}

	return false, nil
}

func (U UserRepository) UpdateUser(db *sql.DB, user model.User) (int64, error) {

	query := "UPDATE users SET username=$1, email=$2, bio=$3, updated_at=$4 WHERE id=$5"

	result, err := db.Exec(query, user.UserName, user.Email, user.Bio, time.Now(), user.Id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil

}
