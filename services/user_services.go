package services

import (
	"errors"
	"library/middleware"
	"library/models"
	"library/repository"
	"library/utils"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) GetUserByID(userIDStr string) (any, any) {
	panic("unimplemented")
}

func (s *UserService) RegisterUser(req *models.User) error {
	// check if user exists in db
	user, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// It's a real DB error, not just "not found"
		return err
	}
	if user != nil {
		return errors.New("user already exists")
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
	token, err := middleware.GenerateJWT(user.ID, user.Role)
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
