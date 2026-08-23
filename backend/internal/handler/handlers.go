package handler

import (
	"backend/internal/database"
	"backend/internal/transform"
	"context"
	"net/http"
)

func (h *Handler) GetEditionResults(w http.ResponseWriter, r *http.Request) {
	matches, ok := h.loadEditionMatches(w, r, h.db.GetEditionResult)

	if ok {
		groupMatchesResponse := transform.GroupByRounds(matches)
		WriteJSON(w, http.StatusOK, groupMatchesResponse)
	}
}

func (h *Handler) GetEditionMatches(w http.ResponseWriter, r *http.Request) {
	matches, ok := h.loadEditionMatches(w, r, h.db.GetEditionMatches)

	if ok {
		groupMatchesResponse := transform.GroupByRounds(matches)
		WriteJSON(w, http.StatusOK, groupMatchesResponse)
	}
}

func (h *Handler) GetFeaturedMatches(w http.ResponseWriter, r *http.Request) {
	matches, ok := h.loadEditionMatches(w, r,
		h.db.FeaturedMatchesQuery(transform.From(), transform.To()),
	)

	if ok {
		featuredMatchesResponse := transform.GetFeaturedMatches(matches)
		WriteJSON(w, http.StatusOK, featuredMatchesResponse)
	}
}

func (h *Handler) loadEditionMatches(w http.ResponseWriter, r *http.Request,
	getEditionMatches func(context.Context, int, int) ([]database.Match, error)) (
	[]database.Match, bool) {

	ctx := r.Context()

	competitionID, startYear, err := editionParams(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return nil, false
	}

	matches, err := getEditionMatches(ctx, competitionID, startYear)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "failed to load matches")
		return nil, false
	}

	return matches, true
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
