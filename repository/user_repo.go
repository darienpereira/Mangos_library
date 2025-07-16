package repository

import (
	"library/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateBook(book models.Book) error
}

type UserRepo struct {
	Db *gorm.DB
}


