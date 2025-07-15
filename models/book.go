package models

import (
	"time"

	"github.com/google/uuid"
)


type Book struct {
	ID uuid.UUID 		`json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title string 		`json:"title"`
	Author string 		`json:"author"`
	Genre string		`json:"genre"`
	Detail string 		`json:"detail" gorm:"type:text"`
	OnShelf bool		`json:"onshelf"`

	Year int				`json:"year"`
	ReturnDate time.Time 	`json:"returndate"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
	DeletedAt time.Time		`json:"deleted_at"`
}

