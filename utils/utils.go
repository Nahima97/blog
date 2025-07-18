package utils

import (
	"golang.org/x/crypto/bcrypt"
)

//hash password with bcrypt
func HashPassword (password string) (string, error) {
hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
if err != nil {
	return "", err
}
return string(hashedPass), nil
}

//compare passwords 
func ComparePassword (hashedPass, plainPass string) error {
err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
if err != nil {
	return err
}
return nil 
}
