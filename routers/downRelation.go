package routers

import (
	"net/http"
	"github.com/giand2205/twittor/db"
	"github.com/giand2205/twittor/models"
)

func DownRelation(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := db.DeletetRelation(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete the relation. Try again "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Could not delete the relation", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}