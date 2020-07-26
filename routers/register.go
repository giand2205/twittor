package routers

import (
	"encoding/json"
	"github.com/giand2205/twittor/bd"
	"github.com/giand2205/twittor/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in the received data"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password >= 6 characters", 400)
		return
	}

	_, found, _ := bd.ValidateUser(t.Email)
	if found == true {
		http.Error(w, "Email is already registered", 400)
		return
	}

	_, status, err := bd.InsertRecord(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert the record"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Could not insert user record", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
