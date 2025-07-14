package handlers

import "library/services"

type UserHandler struct {
	Service *services.UserService
}

/*
register user
login user
user info
*/