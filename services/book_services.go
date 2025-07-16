package services

import (
	"library/models"
	"library/repository"
)

type BookService struct {
	Repo repository.BookRepository
}

func(b BookService) CreateBook(book models.Book) error{
	return b.Repo.CreateBook(&book)
}

func (b BookService) UpdateBook(book models.Book) error {
    return b.Repo.UpdateBook(&book)
}

func (b BookService) DeleteBook(book models.Book) error {
    return b.Repo.DeleteBook(book.ID)
}

