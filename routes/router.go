package routes

import (
	"library/handlers"
	"library/middleware"

	"github.com/gorilla/mux"
)

func SetUpRouter(userHandler *handlers.UserHandler, bookHandler *handlers.BookHandler) *mux.Router {
	r := mux.NewRouter()
	/* public routes */
	

	p := r.PathPrefix("/").Subrouter()
	p.Use(middleware.AuthMiddleware)

	/* user routes */
	p.HandleFunc("/mybooks", bookHandler.ListUserBooks).Methods("GET")

	/* admin routes */

	return r
}