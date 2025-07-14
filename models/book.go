package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID uuid.UUID
	Title string
	Author string
	Genre string
	Detail string
	OnShelf bool
	Year int
	ReturnDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}