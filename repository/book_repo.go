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
