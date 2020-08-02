package routers

import (
	"io"
	"net/http"
	"os"
	"github.com/giand2205/twittor/db"
)

func GetAvatar(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")

	if len(ID)<1 {
		http.Error(w, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil{
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	var OpenFile *os.File
	OpenFile, err = os.Open("uploads/avatars/"+profile.Avatar)
	if err != nil{
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error copying the image", http.StatusBadRequest)
	}

}