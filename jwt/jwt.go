package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/giand2205/twittor/models"
	"time"
)

func GenerateJWT(t models.User) (string, error) {
	key := []byte("PersonalKey_ReactGO_Development")
	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.LastName,
		"birthday":  t.Birthday,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.WebSite,
		"id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
