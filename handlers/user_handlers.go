package handlers

import (
	"encoding/json"

	"library/models"
	"library/services"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}

/*
register user
login user
user info
*/

// register handler
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// collect request details
	var signUp models.User
	err := json.NewDecoder(r.Body).Decode(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call service layer
	err = h.Service.RegisterUser(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signUp)
}

// login handler
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login models.User
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call service layer
	token, err := h.Service.Login(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}
