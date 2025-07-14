package utils

import "golang.org/x/crypto/bcrypt"

type contextKey string

const UserContextKey = contextKey("userClaims")

func HashPassword(p string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(p),bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

func ComparePassword(hashPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	if err != nil {
		return err
	}
	return nil
}