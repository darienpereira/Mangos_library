package repository

import (
	"gorm.io/gorm"
	"library/database"
	"library/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	GetUserById(ID string) (*models.User, error)
}

type UserRepo struct {
	Db *gorm.DB
}


func (r *UserRepo) GetUserByEmail (email string) (*models.User, error) {
	// check if user exists in db
	var user models.User
	err := database.Db.Where("Email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := database.Db.Create(&user).Error
	return err
}

func (r *UserRepo) GetUserById(ID string) (*models.User, error) {
    var user models.User

    err := database.Db.Preload("books").Where("id= ?", ID).First(&user).Error
        return &user, err
}

