package services

import "library/repository"

type UserService struct {
	Repo repository.UserRepository
}