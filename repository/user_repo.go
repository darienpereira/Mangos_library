package repository

import (
	"library/database"
	"library/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error

}

type UserRepo struct {
	
}

func (r *UserRepo) GetUserByEmail (email string) (*models.User, error) {
	// check if user exists in db
	var user models.User
	err := database.Db.Where("email = ?", email).First(&user).Error
	if err == nil {
		return &models.User{}, err
}

return &user, nil

}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := database.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}