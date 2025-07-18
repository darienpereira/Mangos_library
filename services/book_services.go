package services

import (
	"errors"
	"library/models"
	"library/repository"
	"library/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BookService struct {
	Repo repository.BookRepository
}

func (b BookService) CreateBook(book models.Book) error {
	book.OnShelf = true
	book.ReturnDate = nil
	book.UserID = nil
	return b.Repo.CreateBook(&book)
}

func (s *BookService) UpdateBook(req models.Book, id string) error {
	book, err := s.Repo.GetBookByID(id)
	if err != nil {
		return err
	}
	book.Title = req.Title
	book.Author = req.Author
	book.Genre = req.Genre
	book.Year = req.Year
	book.Detail = req.Detail
	return s.Repo.UpdateBook(book)
}

func (s *BookService) DeleteBook(id string) error {
	
	return s.Repo.DeleteBook(id)
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
	err = s.Repo.UpdateBookStock(book)
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
	if book.UserID == nil || *book.UserID != userID {
		return errors.New("book does not belong to you")
	}
	book.OnShelf = true
	book.ReturnDate = nil
	book.UserID = nil
	err = s.Repo.UpdateBookStock(book)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) FindByGenre(genre string) ([]models.Book, error) {
	pattern := utils.CreatePattern(genre)
	books, err := s.Repo.FindByGenre(pattern)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) FindByTitle(title string) ([]models.Book, error) {
	pattern := utils.CreatePattern(title)
	books, err := s.Repo.FindByTitle(pattern)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) FindByAuthor(author string) ([]models.Book, error) {
	pattern := utils.CreatePattern(author)
	books, err := s.Repo.FindByAuthor(pattern)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) FindByYear(year int) ([]models.Book, error) {
	books, err := s.Repo.FindByYear(year)
	if err != nil {
		return nil, err
	}
	return books, nil
}

//service layer
func (s *BookService) GetBook(id string) (*models.Book, error) {
    book, err:=s.Repo.GetBookByID(id)
    if err != nil {
        return &models.Book{}, err
    }
    return book, nil
}

//service layer
func (s *BookService) ListAllBooks() ([]models.Book, error) {
   book, err := s.Repo.GetAllBooks()
   if err != nil {
       return nil, err
   }
   return book, nil
}
