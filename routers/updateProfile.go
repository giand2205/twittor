package routers

import (
	"encoding/json"
	"github.com/giand2205/twittor/db"
	"github.com/giand2205/twittor/models"
	"net/http"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect Data "+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.UpdateRecord(t, IDUser)
	if err != nil {
		http.Error(w, "An error occurred while trying to modify the record. Try again "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Unable to update user information", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
