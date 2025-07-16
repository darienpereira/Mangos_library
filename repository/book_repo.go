package repository

import (
	"library/database"
	"library/models"
)

type BookRepository interface {
	GetBooksByUser(id string) (*[]models.Book, error)
}

type BookRepo struct {
}

func (r *BookRepo) GetBooksByUser(id string) (*[]models.Book, error) {
	var user models.User
	err := database.Db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return &[]models.Book{}, err
	}
	return &user.Books, nil
}