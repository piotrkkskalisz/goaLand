package server

import (
	"backend/internal/database"
	"backend/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter(c *database.Client) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	handler := handler.NewHandler(c)

	//example simply endpoint
	r.Get("/health", handler.GetHealth)

	return r
}
