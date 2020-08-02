package db

import (
	"github.com/giand2205/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := ValidateUser(email)
	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
