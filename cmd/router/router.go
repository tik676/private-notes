package router

import (
	"net/http"
	"private-notes/api/handlers"
	"private-notes/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func InitRoute() http.Handler {
	r := chi.NewRouter()

	r.Post("/registration", handlers.RegisterUserHandler)
	r.Post("/login", handlers.LoginUserHandler)

	r.Group(func(r chi.Router) {
		r.Use(middleware.MiddlewareCheckJWT)
		r.Get("/me", handlers.HandlerMe)
		r.Post("/notes", handlers.CreateNoteHandler)
	})

	return r
}
