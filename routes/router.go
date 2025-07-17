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
	p.HandleFunc("/me", userHandler.GetUserInfo).Methods("GET") 
	p.HandleFunc("/mybooks", bookHandler.ListUserBooks).Methods("GET")

	p.HandleFunc("/search/genre", bookHandler.FindByGenre).Methods("GET")
	p.HandleFunc("/search/title", bookHandler.FindByTitle).Methods("GET")
	p.HandleFunc("/search/author", bookHandler.FindByAuthor).Methods("GET")
	p.HandleFunc("/search/year", bookHandler.FindByYear).Methods("GET")

	p.HandleFunc("/borrow/{id}", bookHandler.BorrowBook).Methods("PUT")
	p.HandleFunc("/return/{id}", bookHandler.ReturnBook).Methods("PUT")

	/* admin routes */
	a := r.PathPrefix("/").Subrouter()
	a.Use(middleware.AuthAdmin)
	a.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	a.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	a.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")

	return r
}
