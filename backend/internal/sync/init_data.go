package sync

import (
	"backend/internal/api"
	"backend/internal/database"
	"backend/internal/utils"
	"fmt"
	"maps"
	"time"
)

const defaultGoalScorerLimit = 10

func intOrZero(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}

func (s *Sync) InitializeData(targets SeasonTargets) error {
	if err := s.initAreas(); err != nil {
		return err
	}

	if err := s.initCompetitions(targets.competitionCodes()); err != nil {
		return err
	}

	for _, target := range targets {
		if err := s.addSeason(target); err != nil {
			return err
		}
	}
	return nil
}

func (s *Sync) addSeason(target SeasonTarget) error {
	competitionID, ok := s.seasons.CompetitionID(target.CompetitionCode)
	if !ok {
		var err error

		competitionID, err = s.initCompetition(target.CompetitionCode)
		if err != nil {
			return err
		}
	}

	season := Season{
		CompetitionID:   competitionID,
		CompetitionCode: target.CompetitionCode,
		StartYear:       target.StartYear,
	}

	if err := s.initEdition(season); err != nil {
		return err
	}

	s.seasons = append(s.seasons, season)

	if err := s.initTeams(season); err != nil {
		return err
	}

	if err := s.initMatches(season); err != nil {
		return err
	}

	if err := s.initGoalScorers(season, defaultGoalScorerLimit); err != nil {
		return err
	}
	return nil
}

func (s *Sync) initAreas() error {
	apiAreas, err := s.apiClient.FetchAreas()
	if err != nil {
		return err
	}

	dbAreas := make([]database.Area, 0, len(apiAreas))

	for _, area := range apiAreas {
		dbAreas = append(dbAreas, database.Area{
			AreaID:    area.ID,
			Name:      area.Name,
			Code:      area.CountryCode,
			IsCountry: area.ParentArea != nil && *area.ParentArea != "World",
		})
		s.areasByName[area.Name] = area.ID
	}

	return s.databaseClient.SaveAreas(dbAreas)
}
func (s *Sync) initCompetition(code string) (int, error) {
	apiCompetition, err := s.apiClient.FetchCompetition(code)
	if err != nil {
		return 0, err
	}

	return apiCompetition.ID, s.databaseClient.SaveCompetitions([]database.Competition{
		{
			CompetitionID:   apiCompetition.ID,
			Name:            apiCompetition.Name,
			Code:            apiCompetition.Code,
			CompetitionType: apiCompetition.Type,
			AreaID:          apiCompetition.Area.ID,
		},
	})
}

func (s *Sync) initCompetitions(codes map[string]struct{}) error {
	apiCompetitions, err := s.apiClient.FetchCompetitions()
	if err != nil {
		return err
	}

	dbCompetitions := make([]database.Competition, 0, len(apiCompetitions))

	missingCodes := maps.Clone(codes)

	for _, competition := range apiCompetitions {
		if _, ok := codes[competition.Code]; ok {
			delete(missingCodes, competition.Code)

			dbCompetitions = append(dbCompetitions, database.Competition{
				CompetitionID:   competition.ID,
				Name:            competition.Name,
				Code:            competition.Code,
				CompetitionType: competition.Type,
				AreaID:          competition.Area.ID,
			})
		}
	}

	if len(missingCodes) > 0 {
		return fmt.Errorf("competitions not found: %v", maps.Keys(missingCodes))
	}

	return s.databaseClient.SaveCompetitions(dbCompetitions)
}

func (s *Sync) initEdition(season Season) error {

	edition, err := s.apiClient.FetchEdition(season.CompetitionCode, season.StartYear)
	if err != nil {
		return err
	}

	startTime, err := time.Parse(api.DateLayout, edition.StartDate)
	if err != nil {
		return err
	}
	endTime, err := time.Parse(api.DateLayout, edition.EndDate)
	if err != nil {
		return err
	}
	dbEdition := database.Edition{
		CompetitionID: season.CompetitionID,
		StartYear:     season.StartYear,
		Status:        utils.EditionStatus(startTime, endTime),
	}

	return s.databaseClient.SaveEdition(dbEdition)
}

func (s *Sync) initTeams(season Season) error {
	apiTeams, err := s.apiClient.FetchTeams(season.CompetitionCode, season.StartYear)
	if err != nil {
		return err
	}

	dbTeams := make([]database.Team, 0, len(apiTeams))

	for _, team := range apiTeams {
		dbTeams = append(dbTeams, database.Team{
			TeamID:    team.ID,
			FullName:  team.Name,
			ShortName: team.ShortName,
			Stadium:   team.Venue,
			Code:      team.TLA,
			Colors:    team.ClubColors,
			AreaID:    team.Area.ID,
		})
	}

	return s.databaseClient.SaveTeams(dbTeams)
}

func (s *Sync) initMatches(season Season) error {
	apiMatches, err := s.apiClient.FetchMatches(season.CompetitionCode, season.StartYear)
	if err != nil {
		return err
	}

	dbMatches := make([]database.Match, 0, len(apiMatches))

	for _, match := range apiMatches {
		startTime, err := time.Parse(time.RFC3339, match.UtcDate)
		if err != nil {
			return err
		}

		dbMatches = append(dbMatches, database.Match{
			MatchID:           match.ID,
			CompetitionID:     season.CompetitionID,
			StartSeasonYear:   season.StartYear,
			HomeTeamID:        match.HomeTeam.ID,
			AwayTeamID:        match.AwayTeam.ID,
			HomeGoals:         match.Score.FullTime.Home,
			AwayGoals:         match.Score.FullTime.Away,
			HalfTimeHomeGoals: match.Score.HalfTime.Home,
			HalfTimeAwayGoals: match.Score.HalfTime.Away,
			Status:            match.Status,
			StartTime:         startTime,

			// TODO
			// StadiumID
		})
	}

	return s.databaseClient.SaveMatches(dbMatches)
}

func (s *Sync) initGoalScorers(season Season, limit int) error {
	apiGoalScorers, err := s.apiClient.FetchGoalScorers(season.CompetitionCode, season.StartYear, limit)
	if err != nil {
		return err
	}

	dbGoalScorers := make([]database.GoalScorer, 0, len(apiGoalScorers))

	for _, scorer := range apiGoalScorers {
		dbGoalScorers = append(dbGoalScorers, database.GoalScorer{
			GoalScorerID:      scorer.Player.ID,
			CompetitionID:     season.CompetitionID,
			StartSeasonYear:   season.StartYear,
			TeamID:            scorer.Team.ID,
			NationalityAreaID: s.areasByName[scorer.Player.Nationality],
			Name:              scorer.Player.Name,
			Goals:             scorer.Goals,
			Assists:           intOrZero(scorer.Assists),
			GoalsFromPenalty:  intOrZero(scorer.Penalties),
		})
	}

	return s.databaseClient.SaveGoalScorers(dbGoalScorers)
}
