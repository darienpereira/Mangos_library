package handlers

import (
	"encoding/json"
	"fmt"
	"library/database"
	"library/models"
	"library/services"
	"log"
	"net/http"

	"library/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type BookHandler struct {
	Service *services.BookService
}

func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "no claims in context", http.StatusInternalServerError)
		return
	}

	err := b.Service.CreateBook(book, claims) //service layer
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("book successfully created")
}

func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	v := mux.Vars(r)
	bookID := v["id"]
	err := b.Service.UpdateBook(book, bookID) //service layer
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("book successfully updated")
}

func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err := b.Service.DeleteBook(book)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("book successfully deleted")
}

func (h *BookHandler) ListUserBooks(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "no claims in context", http.StatusInternalServerError)
		return
	}
	var myBooks []models.Book
	err := h.Service.ListBookByUserID(&myBooks, claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(myBooks)
}

func (h *BookHandler) BorrowBook(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "no claims in context", http.StatusInternalServerError)
		return
	}
	v := mux.Vars(r)
	bookID := v["id"]

	err := h.Service.BorrowBook(bookID, claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("book has been borrowed")
}

func (h *BookHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "no user id in context", http.StatusInternalServerError)
		return
	}
	v := mux.Vars(r)
	bookID := v["id"]

	err := h.Service.ReturnBook(bookID, claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("book has been returned")
}
func CreatePattern(req string) string {
	return fmt.Sprintf("%%%s%%", req)
}

func (h *BookHandler) FindByGenre(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Genre string `json:"genre"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    pattern := CreatePattern(input.Genre)

    var books []models.Book
    err := database.Db.Where("genre LIKE ?", pattern).Find(&books).Error
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(books)
}
func (h *BookHandler) FindByTitle(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Title string `json:"title"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    pattern := CreatePattern(input.Title)

    var books []models.Book
    err := database.Db.Where("title LIKE ?", pattern).Find(&books).Error
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(books)
}
func (h *BookHandler) FindByAuthor(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Author string `json:"author"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    pattern := CreatePattern(input.Author)

    var books []models.Book
    err := database.Db.Where("author LIKE ?", pattern).Find(&books).Error
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(books)
}
func (h *BookHandler) FindByYear(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Year string `json:"year"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    pattern := CreatePattern(input.Year)

    var books []models.Book
    err := database.Db.Where("year LIKE ?", pattern).Find(&books).Error
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(books)
}


/*
Group 1
list All Books - user
get Book By ID - user
list My Books - user           Done
chechout book - user
check in book - user

Group 2
create book - admin
delete book - admin
update book - admin


Group 3
find book by genre
find book by author
find book by year
find book by title


Group 4
register user
login user
user info
*/
