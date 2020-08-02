package routers

import (
	"encoding/json"
	"github.com/giand2205/twittor/db"
	"net/http"
	"strconv"
)

func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the page parameter", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send only numbers greater than zero in the page parameter", http.StatusBadRequest)
		return
	}

	response, right := db.ReadTweetsFollowers(IDUser, page)
	if right == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
