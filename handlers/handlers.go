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
	router.HandleFunc("/delete-tweet", middlew.ValidateDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/upload-avatar", middlew.ValidateDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/get-avatar", middlew.ValidateDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/upload-banner", middlew.ValidateDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/get-banner", middlew.ValidateDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/up-relation", middlew.ValidateDB(middlew.ValidateJWT(routers.UpRelation))).Methods("POST")
	router.HandleFunc("/down-relation", middlew.ValidateDB(middlew.ValidateJWT(routers.DownRelation))).Methods("DELETE")
	router.HandleFunc("/consult-relation", middlew.ValidateDB(middlew.ValidateJWT(routers.ConsultRelation))).Methods("GET")
	router.HandleFunc("/list-users", middlew.ValidateDB(middlew.ValidateJWT(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/read-tweets-followers", middlew.ValidateDB(middlew.ValidateJWT(routers.ReadTweetsFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
