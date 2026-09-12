package handler

import (
	"backend/internal/database"
	"backend/internal/sync"
	"backend/internal/transform"
	"backend/internal/utils"
	"context"
	"net/http"
	"slices"
	"time"
)

const (
	UpcomingStatus = "upcoming"
	LiveStatus     = "live"
	FinishedStatus = "finished"
)

const maxMatchesOnTop = 50

func (h *Handler) GetAllMatches(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	startYear := sync.CurrentSeasonStartYear(time.Now())

	var allMatches [][]transform.MatchResponse
	for _, statuses := range utils.DisplayableMatchStatusesLists {
		matches, err := h.db.GetAllMatchesWithStatus(ctx, startYear, statuses)
		if err != nil {
			WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		if slices.Compare(statuses, utils.LiveMatchStatuses) == 0 {
			allMatches = append(allMatches,
				transform.GetFeaturedMatches(matches))
		} else {
			allMatches = append(allMatches,
				transform.GetResponseMatchesWithLimit(matches, maxMatchesOnTop))
		}
	}

	WriteJSON(w, http.StatusOK, allMatches)
}

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

func (h *Handler) GetClubPlayers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	teamID, err := teamParam(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	startYear := sync.CurrentSeasonStartYear(time.Now())

	competitionID, err := competitionParam(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	var players []database.SeasonPlayer
	if err := h.db.List(ctx, &players, database.Filter{
		"competition_id": competitionID,
		"start_year":     startYear,
		"team_id":        teamID,
	}, "Player"); err != nil {
		WriteError(w, http.StatusInternalServerError, "failed to load club players")
		return
	}

	WriteJSON(w, http.StatusOK, players)
}

func (h *Handler) GetTeamInformation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	clubID, err := clubParam(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	team, err := h.db.GetTeam(ctx, clubID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "failed to load club information")
		return
	}

	WriteJSON(w, http.StatusOK, team)
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

	players, err := h.db.GetEditionGoalScorers(ctx, competitionID, startYear)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "failed to load goal scorers")
		return
	}

	response := transform.GetEditionGoalScorers(players)

	WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) GetEditionTable(w http.ResponseWriter, r *http.Request) {
	competitionID, startYear, err := editionParams(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	clubs, ok := h.store.GetTableClubs(competitionID, startYear)
	if !ok {
		WriteError(w, http.StatusInternalServerError, "failed to load table")
		return
	}

	WriteJSON(w, http.StatusOK, clubs)

}
