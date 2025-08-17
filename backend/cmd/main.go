package main

import (
	"log"
	"net/http"
	"private-notes/cmd/router"
	"private-notes/internal/db"
)

func main() {
	db.InitDB()

	r := router.InitRoute()
	go db.RegularClearNoteByExpires()

	log.Println("Server is starting on port :2288")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server startup error: %v", err)
		return
	}
}
