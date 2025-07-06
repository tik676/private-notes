package router

import (
	"net/http"
	"private-notes/api/handlers"
	"private-notes/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func InitRoute() {
	r := chi.NewRouter()

	r.Use(middleware.MiddlewareCheckJWT)
	r.Get("/me", handlers.HandlerMe)

	http.ListenAndServe(":2288", r)
}
