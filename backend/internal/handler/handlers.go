package handler

import (
	"backend/internal/database"
	"backend/internal/transform"
	"net/http"
)

func (h *Handler) GetEditionMatches(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	competitionID, startYear, err := editionParams(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.db.GetEditionMatches(ctx, competitionID, startYear)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "failed to load matches")
		return
	}

	WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) GetTeamsMatches(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	teamsID, err := teamParam(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.db.GetTeamsMatches(ctx, teamsID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "failed to load matches")
		return
	}

	WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) GetCompetitionEdition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var competitions []database.Competition
	if err := h.db.List(ctx, &competitions, nil, database.PreloadEditions); err != nil {
		http.Error(w, "failed to load leagues", http.StatusInternalServerError)
		return
	}

	response := transform.GetCompetitionEdition(competitions)

	WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) GetEditionGoalScorers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	competitionID, startYear, err := editionParams(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.db.GetEditionGoalScorers(ctx, competitionID, startYear)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "failed to load goal scorers")
		return
	}

	WriteJSON(w, http.StatusOK, response)
}
