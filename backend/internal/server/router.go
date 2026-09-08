package server

import (
	"backend/internal/database"
	"backend/internal/handler"
	"backend/internal/state"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(db *database.Client, store *state.Store) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	handler := handler.NewHandler(db, store)

	//example simply endpoint
	r.Get("/health", handler.GetHealth)

	// Endpoints
	r.Get("/teams/{teamID}/matches", handler.GetTeamsMatches)
	r.Get("/competitions", handler.GetCompetitionEdition)
	r.Get("/competitions/{competitionID}/{startYear}/matches", handler.GetEditionMatches)
	r.Get("/competitions/{competitionID}/{startYear}/goal-scorers", handler.GetEditionGoalScorers)
	r.Get("/competitions/{competitionID}/{startYear}/results", handler.GetEditionResults)
	r.Get("/competitions/{competitionID}/{startYear}/featured-matches", handler.GetFeaturedMatches)

	return r
}
