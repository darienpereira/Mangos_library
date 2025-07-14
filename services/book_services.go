package services

import "library/repository"

type BookService struct {
	Repo repository.BookRepository
}