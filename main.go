package main

import (
	"github.com/giand2205/twittor/bd"
	"github.com/giand2205/twittor/handlers"
	"log"
)

func main() {
	if bd.ValidateConnection() == 0 {
		log.Fatal("No connection to DB")
		return
	}
	handlers.Handlers()
}
