package repository

import (
	"library/database"
	"library/models"

	"github.com/google/uuid"
)

type BookRepository interface {
    CreateBook(book *models.Book) error
    UpdateBook(book *models.Book) error
    DeleteBook(id uuid.UUID) error
)

type BookRepository interface {
	GetBooksByUser(id string) (*[]models.Book, error)
	GetBookByID(id string) (*models.Book, error)
	UpdateBook(update *models.Book) error
}

type BookRepo struct {
}

func (r *BookRepo) CreateBook(book *models.Book) error {
    return database.Db.Create(book).Error
}

func (r *BookRepo) UpdateBook(book *models.Book) error {
    return database.Db.Model(&models.Book{}).Where("id = ?", book.ID).Updates(book).Error
}

func (r *BookRepo) DeleteBook(id uint) error {
    return database.Db.Delete(&models.Book{}, id).Error
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

