package handlers

import (
	"encoding/json"

	"library/models"
	"library/services"
	"library/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call service layer
	err = h.Service.RegisterUser(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
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
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}

func (h *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	claims,ok:=r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if	!ok {
		http.Error(w, "unable to get claim", http.StatusInternalServerError)
		return
	}
	user, err := h.Service.GetUserInfo(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
