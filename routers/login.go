package routers

import (
	"encoding/json"
	"github.com/giand2205/twittor/bd"
	"github.com/giand2205/twittor/jwt"
	"github.com/giand2205/twittor/models"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid user or password "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	document, exist := bd.TryLogin(t.Email, t.Password)
	if exist == false {
		http.Error(w, "Invalid user or password", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error generating Token", 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}
