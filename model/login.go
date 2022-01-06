package model

import "errors"

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (loginDTO *LoginDTO) Validate() error {

	if len(loginDTO.Email) == 0 {
		return errors.New("email should not be empty")
	}

	if len(loginDTO.Password) == 0 {
		return errors.New("password should not be empty")
	}

	return nil
}
