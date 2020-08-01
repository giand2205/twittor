package routers

import (
	"encoding/json"
	"github.com/giand2205/twittor/bd"
	"net/http"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "An error occurred while trying to find the record "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
