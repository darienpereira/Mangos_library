package repository

import (
	"errors"
	"library/database"
	"library/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	GetUserById(ID string) (*models.User, error)
}

type UserRepo struct {
	Db *gorm.DB
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.Db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // user does not exist
	}
	if err != nil {
		return nil, err // real DB error
	}
	return &user, nil // user exists
}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := database.Db.Create(&user).Error
	return err
}

func (r *UserRepo) GetUserById(ID string) (*models.User, error) {
	var user models.User

	err := database.Db.Preload("Books").Where("id= ?", ID).First(&user).Error
	return &user, err
}
