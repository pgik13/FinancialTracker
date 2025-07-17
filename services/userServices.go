package services

import (
	"errors"
	"trackytrack/middleware"
	"trackytrack/models"
	"trackytrack/repo"
	"trackytrack/utils"
)

type UserServices struct {
	Repo repo.UserRepo
}

func (s *UserServices) RegisterUser(user *models.User) error {
	//check if user exists
	existingUser, err := s.Repo.GetUserByEmail(user.Email)

	//DB error
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("this email already has an account")
	}

	hashpass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashpass

	err = s.Repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServices) LogInUser(request models.LoginRequest) (string, error) {
	user, err := s.Repo.GetUserByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, request.Password)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
