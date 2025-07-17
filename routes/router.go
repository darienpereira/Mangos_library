package routes

import (
	"library/handlers"
	"library/middleware"

	"github.com/gorilla/mux"
)

func SetUpRouter(userHandler *handlers.UserHandler, bookHandler *handlers.BookHandler) *mux.Router {
	r := mux.NewRouter()

	/* public routes */
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")

	p := r.PathPrefix("/").Subrouter()
	p.Use(middleware.AuthMiddleware)

	/* user routes */
	p.HandleFunc("/mybooks", bookHandler.ListUserBooks).Methods("GET")

	p.HandleFunc("/search/genre", bookHandler.FindByGenre).Methods("POST")
	p.HandleFunc("/search/title", bookHandler.FindByTitle).Methods("POST")
	p.HandleFunc("/search/author", bookHandler.FindByAuthor).Methods("POST")
	p.HandleFunc("/search/year", bookHandler.FindByYear).Methods("POST")

	/* admin routes */

	return r
}
