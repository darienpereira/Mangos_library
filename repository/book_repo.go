package repository

import (
	"library/database"
	"library/models"
)

type BookRepository interface {
	GetBooksByUser(id string) (*[]models.Book, error)
	GetBookByID(id string) (*models.Book, error)
	UpdateBook(update *models.Book) error
}

type BookRepo struct {
}

func (r *BookRepo) GetBooksByUser(id string) (*[]models.Book, error) {
	var user models.User
	err := database.Db.Where("ID = ?", id).First(&user).Error
	return &user.Books, err
}

func (r *BookRepo) GetBookByID(id string) (*models.Book, error) {
	var book models.Book
	err := database.Db.Where("ID = ?", id).First(&book).Error
	return &book, err
}

func (r *BookRepo) UpdateBook(update *models.Book) error {
	err := database.Db.Save(update).Error
	return err
}