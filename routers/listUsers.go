package routers

import (
	"encoding/json"
	"github.com/giand2205/twittor/db"
	"net/http"
	"strconv"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeSearch := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "You must send only numbers greater than zero in the page parameter", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := db.ReadAllUsers(IDUser, pag, search, typeSearch)

	if status == false {
		http.Error(w, "Error reading users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
