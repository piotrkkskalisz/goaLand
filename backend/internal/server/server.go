package server

import (
	"backend/internal/database"
	"backend/internal/state"
	"net/http"
)

const BackendPort = ":8080"

type Server struct {
	*http.Server
}

func NewServer(db *database.Client, store *state.Store) *Server {
	router := NewRouter(db, store)

	return &Server{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
}
