package services

import (
	"errors"
	"library/models"
	"library/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BookService struct {
	Repo repository.BookRepository
}

func (b *BookService) FindByAuthor(pattern string) (any, any) {
	panic("unimplemented")
}

func (b BookService) CreateBook(book models.Book) error {
	return b.Repo.CreateBook(&book)
}

func (b BookService) UpdateBook(book models.Book) error {
	return b.Repo.UpdateBook(&book)
}

func (b BookService) DeleteBook(book models.Book) error {
	return b.Repo.DeleteBook(book.ID)
}

func (s *BookService) ListBookByUserID(books *[]models.Book, claims jwt.MapClaims) error {
	userId, ok := claims["userID"].(string)
	if !ok {
		return errors.New("no userid in claims")
	}
	books, err := s.Repo.GetBooksByUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) BorrowBook(bookID string, claims jwt.MapClaims) error {
	userID, err := uuid.Parse(claims["userID"].(string))
	if err != nil {
		return err
	}
	book, err := s.Repo.GetBookByID(bookID)
	if err != nil {
		return err
	}
	if !book.OnShelf {
		return errors.New("book is not in stock")
	}
	book.OnShelf = false
	returnDate := time.Now().Add(24 * time.Hour)
	book.ReturnDate = &returnDate
	book.UserID = &userID
	err = s.Repo.UpdateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) ReturnBook(bookID string, claims jwt.MapClaims) error {
	userID, err := uuid.Parse(claims["userID"].(string))
	if err != nil {
		return err
	}
	book, err := s.Repo.GetBookByID(bookID)
	if err != nil {
		return err
	}
	if book.UserID != &userID {
		return errors.New("book belongs to someone else")
	}
	book.OnShelf = true
	book.ReturnDate = nil
	book.UserID = nil
	err = s.Repo.UpdateBook(book)
	if err != nil {
		return err
	}

	return nil
}
