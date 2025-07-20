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

	log.Println("🔥 Сервер запускается на порту :2288")

	err := http.ListenAndServe(":2288", r)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
		return
	}

}
