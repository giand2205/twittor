package main

import (
	"github.com/giand2205/twittor/db"
	"github.com/giand2205/twittor/handlers"
	"log"
)

func main() {
	if db.ValidateConnection() == 0 {
		log.Fatal("No connection to DB")
		return
	}
	handlers.Handlers()
}
