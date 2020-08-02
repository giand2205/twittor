package routers

import (
	"encoding/json"
	"github.com/giand2205/twittor/db"
	"github.com/giand2205/twittor/models"
	"net/http"
)

func ConsultRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	var resp models.ResponseConsultRelation

	status, err := db.ConsultRelation(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
