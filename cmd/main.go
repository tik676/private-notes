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
	err := http.ListenAndServe(":2288", r)
	if err != nil {
		log.Fatal("жопа")
		return
	}

}
