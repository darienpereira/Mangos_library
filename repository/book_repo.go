package repository

import (
	"library/database"
	"library/models"
)

type BookRepository interface {
	GetBooksByUser(id string) (*[]models.Book, error)
	GetBookByID(id string) (*models.Book, error)
	UpdateBook(update *models.Book) error
	CreateBook(book *models.Book) error
    DeleteBook(id string) error
	UpdateBookStock(update *models.Book) error
	FindByYear(pattern string) ([]models.Book, error)
	FindByAuthor(pattern string) ([]models.Book, error)
	FindByTitle(pattern string) ([]models.Book, error)
	FindByGenre(pattern string) ([]models.Book, error)
}

type BookRepo struct {
}

func (r *BookRepo) CreateBook(book *models.Book) error {
    return database.Db.Create(book).Error
}

func (r *BookRepo) UpdateBook(book *models.Book) error {
    return database.Db.Model(&models.Book{}).Where("id = ?", book.ID).Updates(book).Error
}

func (r *BookRepo) DeleteBook(id string) error {
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

func (r *BookRepo) UpdateBookStock(update *models.Book) error {
	err := database.Db.Save(update).Error
	return err
}

func (r *BookRepo) FindByGenre(pattern string) ([]models.Book, error) {
	var books []models.Book
	err := database.Db.Where("genre LIKE ?", pattern).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepo) FindByTitle(pattern string) ([]models.Book, error) {
	var books []models.Book
	err := database.Db.Where("title LIKE ?", pattern).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepo) FindByAuthor(pattern string) ([]models.Book, error) {
	var books []models.Book
	err := database.Db.Where("author LIKE ?", pattern).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepo) FindByYear(pattern string) ([]models.Book, error) {
	var books []models.Book
	err := database.Db.Where("CAST(year AS TEXT) LIKE ?", pattern).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

