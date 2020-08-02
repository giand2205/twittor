package routers

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/giand2205/twittor/db"
	"github.com/giand2205/twittor/models"
	"strings"
)

var Email string
var IDUser string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	key := []byte("PersonalKey_ReactGO_Development")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid Token Format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		_, found, _ := db.ValidateUser(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid Token")
	}
	return claims, false, string(""), err
}
