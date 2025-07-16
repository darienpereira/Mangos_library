package services

import (
	"library/middleware"
	"library/models"
	"library/repository"
	"library/utils"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) RegisterUser(req *models.User) error {
	// check if user exists in db
	_, err := s.Repo.GetUserByEmail(req.Email)
	if err == nil {
		return err
	}

	// hash the password
	HashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = HashPassword

	// add user into db
	err = s.Repo.CreateUser(req)
	if err != nil {
		return err
	}
	return nil

}

func (s *UserService) Login(req *models.User) (string, error) {
	//check if user exists
	user, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		return "", err
	}

	//generate token
	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserService) GetUserInfo(claims jwt.MapClaims) (*models.User, error) {
	id := claims["userID"].(string)
	user, err := s.Repo.GetUserById(id)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}
