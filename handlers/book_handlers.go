package handlers

import (
	"encoding/json"
	"library/models"
	"library/services"
	"library/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type BookHandler struct {
	Service *services.BookService
}

func (h *BookHandler) ListUserBooks(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "no user id in context", http.StatusInternalServerError)
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