package router

import (
	"net/http"
	"private-notes/api/handlers"
	"private-notes/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func InitRoute() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.CORSMiddleware)

	r.Post("/register", handlers.RegisterUserHandler)
	r.Post("/login", handlers.LoginUserHandler)
	r.Post("/refresh-token", handlers.RefreshTokenHandle)
	r.Post("/logout", handlers.LogoutHandler)

	r.Get("/notes/public/{id}", handlers.GetPublicNoteHandler)

	r.Group(func(r chi.Router) {
		r.Use(middleware.MiddlewareCheckJWT)

		r.Get("/me", handlers.HandlerMe)

		r.Post("/notes", handlers.CreateNoteHandler)

		r.Patch("/notes/{id}", handlers.UpdateNoteHandler)

		r.Delete("/notes/{id}", handlers.DeleteNoteHandler)
	})

	return r
}
