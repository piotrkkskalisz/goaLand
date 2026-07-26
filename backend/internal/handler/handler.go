package handler

import (
	"backend/internal/database"
	"net/http"
)

type Handler struct {
	db *database.Client
}

func NewHandler(db *database.Client) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Ok"))
}
