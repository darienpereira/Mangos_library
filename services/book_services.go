package services

import (
	"errors"
	"library/models"
	"library/repository"

	"github.com/golang-jwt/jwt/v5"
)

type BookService struct {
	Repo repository.BookRepository
}

func (s *BookService) ListBookByUserID(books *[]models.Book, claims jwt.MapClaims) error{
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