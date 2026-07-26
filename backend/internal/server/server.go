package server

import (
	"backend/internal/database"
	"net/http"
)

const BackendPort = ":8080"

type Server struct {
	*http.Server
}

func NewServer(c *database.Client) *Server {
	router := NewRouter(c)

	return &Server{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
}
