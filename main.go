package main

import (
	"fmt"
	"library/database"
	"library/handlers"
	"library/repository"
	"library/routes"
	"library/services"
	"library/utils"
	"log"
	"net/http"
)

func main() {
	/* initialise database */
	db := database.OpenDb()

	/* initialise repositories */
	userRepo := &repository.UserRepo{}
	bookRepo := &repository.BookRepo{}

	/* initialise service */
	userService := &services.UserService{Repo: userRepo}
	bookService := &services.BookService{Repo: bookRepo}

	/* initialise handlers */
	userHandler := &handlers.UserHandler{Service: userService}
	bookHandler := &handlers.BookHandler{Service: bookService}

	/* initialise routes */
	r := routes.SetUpRouter(userHandler, bookHandler)

	/* start server */
	fmt.Println("starting server...")
	err := http.ListenAndServe(":"+utils.GetPort(), r)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}

/*
Imran            Darien    Sabir 1
Sevi             Jasmine         2
Valencia         Jordan          3
Nourhan          Zaahirah        4
*/