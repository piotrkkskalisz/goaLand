package handler

import (
	"backend/internal/database"
	"backend/internal/transform"
	"context"
	"net/http"
)

func (h *Handler) GetEditionResults(w http.ResponseWriter, r *http.Request) {
	h.GetEditionSelectMatches(w, r, h.db.GetEditionResult)
}

func (h *Handler) GetEditionMatches(w http.ResponseWriter, r *http.Request) {
	h.GetEditionSelectMatches(w, r, h.db.GetEditionMatches)
}
func (h *Handler) GetEditionSelectMatches(w http.ResponseWriter, r *http.Request,
	getEditionMatches func(context.Context, int, int) ([]database.Match, error)) {

	ctx := r.Context()

	competitionID, startYear, err := editionParams(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	matches, err := getEditionMatches(ctx, competitionID, startYear)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "failed to load matches")
		return
	}

	groupMatchesResponse := transform.GroupByRounds(matches)

	WriteJSON(w, http.StatusOK, groupMatchesResponse)
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
