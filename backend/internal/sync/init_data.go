package sync

import (
	"backend/internal/api"
	"backend/internal/database"
	"backend/internal/utils"
	"context"
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

func (s *Sync) InitializeData(ctx context.Context, targets SeasonTargets) error {
	now := time.Now()
	if err := s.initAreas(ctx); err != nil {
		return err
	}

	if err := s.initCompetitions(ctx, targets.competitionCodes()); err != nil {
		return err
	}

	for _, target := range targets {
		if err := s.addSeason(ctx, now, target); err != nil {
			return err
		}
	}
	return nil
}

func (s *Sync) addSeason(ctx context.Context, now time.Time, target SeasonTarget) error {
	competitionID, ok := s.seasons.CompetitionID(target.CompetitionCode)
	if !ok {
		var err error

		competitionID, err = s.initCompetition(ctx, target.CompetitionCode)
		if err != nil {
			return err
		}
	}

	season := Season{
		CompetitionID:   competitionID,
		CompetitionCode: target.CompetitionCode,
		StartYear:       target.StartYear,
	}

	if err := s.initEdition(ctx, season); err != nil {
		return err
	}

	s.seasons = append(s.seasons, season)

	if err := s.initTeams(ctx, now, season); err != nil {
		return err
	}

	if err := s.initMatches(ctx, season); err != nil {
		return err
	}

	if err := s.initGoalScorers(ctx, season, defaultGoalScorerLimit); err != nil {
		return err
	}
	return nil
}

func (s *Sync) initAreas(ctx context.Context) error {
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

	return s.databaseClient.Save(ctx, dbAreas)
}
func (s *Sync) initCompetition(ctx context.Context, code string) (int, error) {
	apiCompetition, err := s.apiClient.FetchCompetition(code)
	if err != nil {
		return 0, err
	}

	return apiCompetition.ID, s.databaseClient.Save(ctx, []database.Competition{
		{
			CompetitionID:   apiCompetition.ID,
			Name:            apiCompetition.Name,
			Code:            apiCompetition.Code,
			CompetitionType: apiCompetition.Type,
			AreaID:          apiCompetition.Area.ID,
		},
	})
}

func (s *Sync) initCompetitions(ctx context.Context, codes map[string]struct{}) error {
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

	return s.databaseClient.Save(ctx, dbCompetitions)
}

func (s *Sync) initEdition(ctx context.Context, season Season) error {

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

	return s.databaseClient.Save(ctx, &dbEdition)
}

func (s *Sync) initPlayers(ctx context.Context, apiTeams []api.Team, season Season) error {
	dbPlayers := make([]database.Player, 0, 0)
	dbSeasonPlayers := make([]database.SeasonPlayer, 0, 0)

	for _, team := range apiTeams {
		for _, player := range team.Players {
			if _, exist := s.areasByName[player.Nationality]; !exist {
				return fmt.Errorf("Not found Area %s", player.Nationality)
			}

			dbPlayers = append(dbPlayers, database.Player{
				PlayerID:          player.ID,
				Name:              player.Name,
				NationalityAreaID: s.areasByName[player.Nationality],
				Position:          player.Position,
			})
			dbSeasonPlayers = append(dbSeasonPlayers, database.SeasonPlayer{
				PlayerID:        player.ID,
				TeamID:          team.ID,
				CompetitionID:   season.CompetitionID,
				StartSeasonYear: season.StartYear,
			})

		}
	}
	if err := s.databaseClient.Save(ctx, dbPlayers); err != nil {
		return err
	}
	return s.databaseClient.Save(ctx, dbSeasonPlayers)
}

func (s *Sync) initTeams(ctx context.Context, now time.Time, season Season) error {
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

	if err := s.databaseClient.Save(ctx, dbTeams); err != nil {
		return err
	}

	if isCurrentSeason(season.StartYear, now) {
		if err = s.initPlayers(ctx, apiTeams, season); err != nil {
			return err
		}
	}
	return nil

}

func (s *Sync) initMatches(ctx context.Context, season Season) error {
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

			Matchday: match.Matchday,
			Stage:    match.Stage,

			// TODO
			// StadiumID
		})
	}

	return s.databaseClient.Save(ctx, dbMatches)
}

func (s *Sync) initGoalScorers(ctx context.Context, season Season, limit int) error {
	apiGoalScorers, err := s.apiClient.FetchGoalScorers(season.CompetitionCode, season.StartYear, limit)
	if err != nil {
		return err
	}

	dbPlayers := make([]database.Player, 0, len(apiGoalScorers))
	dbSeasonPlayers := make([]database.SeasonPlayer, 0, len(apiGoalScorers))

	for _, scorer := range apiGoalScorers {
		areaID, exists := s.areasByName[scorer.Player.Nationality]
		if !exists {
			return fmt.Errorf("not found area %s", scorer.Player.Nationality)
		}

		dbPlayers = append(dbPlayers, database.Player{
			PlayerID:          scorer.Player.ID,
			Name:              scorer.Player.Name,
			Position:          scorer.Player.Section,
			NationalityAreaID: areaID,
		})

		dbSeasonPlayers = append(dbSeasonPlayers, database.SeasonPlayer{
			PlayerID:         scorer.Player.ID,
			CompetitionID:    season.CompetitionID,
			StartSeasonYear:  season.StartYear,
			TeamID:           scorer.Team.ID,
			Goals:            scorer.Goals,
			Assists:          scorer.Assists,
			GoalsFromPenalty: scorer.Penalties,
		})
	}

	if len(apiGoalScorers) == 0 {
		return nil
	}

	if err := s.databaseClient.Save(ctx, dbPlayers); err != nil {
		return err
	}

	return s.databaseClient.Save(ctx, dbSeasonPlayers)
}
