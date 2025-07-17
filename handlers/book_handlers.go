package handlers

import (
	"encoding/json"
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

	err := b.Service.CreateBook(book) //service layer
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("book successfully updated")
}

func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	bookID := v["id"]

	err := b.Service.DeleteBook(bookID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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

func (h *BookHandler) FindByGenre(w http.ResponseWriter, r *http.Request) {
	var input models.BookSearch

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	books, err := h.Service.FindByGenre(input.Genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) FindByTitle(w http.ResponseWriter, r *http.Request) {
	var input models.BookSearch

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	books, err := h.Service.FindByTitle(input.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) FindByAuthor(w http.ResponseWriter, r *http.Request) {
	var input models.BookSearch

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	books, err := h.Service.FindByAuthor(input.Author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) FindByYear(w http.ResponseWriter, r *http.Request) {
	var input models.BookSearch

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	books, err := h.Service.FindByYear(input.Year) // ✅ Call the service, not database directly
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
