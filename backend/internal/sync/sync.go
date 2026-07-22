package sync

import (
	"backend/internal/api"
	"backend/internal/database"
)

type SeasonTarget struct {
	CompetitionCode string
	StartYear       int
}

type SeasonTargets []SeasonTarget

type Season struct {
	CompetitionID   int
	CompetitionCode string
	StartYear       int
}
type Seasons []Season

// TODO: load seasons and areas from DB on start working if DB is not empty

//go:generate mockgen -source=sync.go -destination=mocks/sync_mock.go -package=mocks
type API interface {
	FetchAreas() ([]api.Area, error)
	FetchCompetition(code string) (api.Competition, error)
	FetchCompetitions() ([]api.Competition, error)

	FetchEdition(code string, startYear int) (api.Edition, error)
	FetchTeams(code string, startYear int) ([]api.Team, error)
	FetchMatches(code string, startYear int) ([]api.Match, error)
	FetchGoalScorers(code string, startYear, limit int) ([]api.GoalScorer, error)
}

type Database interface {
	SaveAreas([]database.Area) error
	SaveCompetitions([]database.Competition) error
	SaveEdition(database.Edition) error
	SaveTeams([]database.Team) error
	SaveMatches([]database.Match) error
	SaveGoalScorers([]database.GoalScorer) error
}

type Sync struct {
	apiClient      API
	databaseClient Database
	seasons        Seasons
	areasByName    map[string]int
}

func New(apiClient API, databaseClient Database) *Sync {
	return &Sync{
		apiClient:      apiClient,
		databaseClient: databaseClient,
		areasByName:    make(map[string]int),
	}
}

func NewFromEnv() (*Sync, error) {
	databaseClient, err := database.NewClientFromEnv()
	if err != nil {
		return nil, err
	}
	return &Sync{
		apiClient:      api.NewClientFromEnv(),
		databaseClient: databaseClient,
		seasons:        nil,
		areasByName:    make(map[string]int),
	}, nil
}

func (targets SeasonTargets) competitionCodes() map[string]struct{} {
	set := make(map[string]struct{})

	for _, target := range targets {
		set[target.CompetitionCode] = struct{}{}
	}

	return set
}

// #TODO reduce complexity from O(n)
func (seasons Seasons) CompetitionID(code string) (int, bool) {
	for _, season := range seasons {
		if season.CompetitionCode == code {
			return season.CompetitionID, true
		}
	}

	return 0, false
}
