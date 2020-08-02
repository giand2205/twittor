package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/giand2205/twittor/bd"
	"github.com/giand2205/twittor/models"
)

func SaveTweet(w http.ResponseWriter, r *http.Request){
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	record := models.SaveTweet{
		UserID: IDUser,
		Message: message.Message,
		Date: time.Now(),
	}

	var status bool
	_, status, err = bd.InsertTweet(record)

	if err != nil {
		http.Error(w, "An error occurred while trying to insert the record. Try again "+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "Unable to insert the Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}