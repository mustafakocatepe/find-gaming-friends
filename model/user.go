package model

import (
	"errors"
	"time"
)

// User is our main model for Users
type User struct {
	Id        int       `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Bio       string    `json:"bio"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

// UserDTO is our data transfer object for User
type UserDTO struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type CreateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

func (user *CreateUserDTO) Validate() error {

	if len(user.Username) == 0 {
		return errors.New("username should not be empty")
	}

	if len(user.Email) == 0 {
		return errors.New("email should not be empty")
	}

	if len(user.Password) == 0 {
		return errors.New("password should not be empty")
	}

	return nil
}
