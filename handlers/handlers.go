package handlers

import (
	"github.com/giand2205/twittor/middlew"
	"github.com/giand2205/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.ValidateDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.ValidateDB(routers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middlew.ValidateDB(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/update-profile", middlew.ValidateDB(middlew.ValidateJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ValidateDB(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/read-tweets", middlew.ValidateDB(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
