package testutils

import (
	"backend/internal/database"
	"time"
)

const EnglandName = "England"

const Year = 2025

const EnglandAreaID = 1
const PolandAreaID = 2

const PremierLeagueCode = "PL"
const LaLigaCode = "PD"

const PremierLeagueID = 2021
const EkstraklasaID = 1950

const ArsenalID = 24
const ManCityID = 25
const ChelseaId = 26
const LiverpoolID = 27

const SakaID = 2

// NewArea creates England area.
func NewArea() database.Area {
	return England()
}

func England() database.Area {
	return database.Area{
		AreaID:    EnglandAreaID,
		Name:      "England",
		Code:      "EN",
		IsCountry: true,
	}
}

func Poland() database.Area {
	return database.Area{
		AreaID:    PolandAreaID,
		Name:      "Poland",
		Code:      "PL",
		IsCountry: true,
	}
}

// NewCompetition creates Premier League competition.
func NewCompetition() database.Competition {
	return PremierLeague()
}

func Ekstraklasa() database.Competition {
	return database.Competition{
		CompetitionID:   EkstraklasaID,
		Name:            "Ekstraklasa",
		Code:            "EK",
		CompetitionType: "LEAGUE",
		AreaID:          PolandAreaID,
	}
}

func PremierLeague() database.Competition {
	return database.Competition{
		CompetitionID:   PremierLeagueID,
		Name:            "Premier League",
		Code:            PremierLeagueCode,
		CompetitionType: "LEAGUE",
		AreaID:          EnglandAreaID,
	}
}

// NewTeam creates Arsenal FC team.
func NewTeam() database.Team {
	return Arsenal()
}

func Arsenal() database.Team {
	return database.Team{
		TeamID:    ArsenalID,
		FullName:  "Arsenal FC",
		ShortName: "Arsenal",
		Code:      "ARS",
		Colors:    "Red / White",
		AreaID:    EnglandAreaID,
		Stadium:   "Emirates Stadium",
	}
}

func Chelsea() database.Team {
	return database.Team{
		TeamID:    ChelseaId,
		FullName:  "Chelsea FC",
		ShortName: "Chelsea",
		Code:      "CHE",
		Colors:    "Blue / White",
		AreaID:    EnglandAreaID,
		Stadium:   "Stamford Bridge",
	}
}

func ManCity() database.Team {
	return database.Team{
		TeamID:    ManCityID,
		FullName:  "Manchester City FC",
		ShortName: "Man City",
		Code:      "MCI",
		Colors:    "Blue",
		AreaID:    EnglandAreaID,
	}
}

func Liverpool() database.Team {
	return database.Team{
		TeamID:    LiverpoolID,
		FullName:  "Liverpool FC",
		ShortName: "Liverpool",
		Code:      "LIV",
		Colors:    "Red",
		AreaID:    EnglandAreaID,
		Stadium:   "Anfield",
	}
}

// NewEdition creates 2025 Premier League edition.
func NewEdition() database.Edition {
	return database.Edition{
		CompetitionID: PremierLeagueID,
		StartYear:     Year,
		Status:        "FINISHED",
	}
}

// NewMatch creates Arsenal vs Manchester City match.
func NewMatch() database.Match {
	homeGoals, awayGoals := 2, 1
	matchday := 1

	return database.Match{
		MatchID:         100,
		CompetitionID:   PremierLeagueID,
		StartSeasonYear: Year,
		HomeTeamID:      ArsenalID,
		AwayTeamID:      ManCityID,
		HomeGoals:       &homeGoals,
		AwayGoals:       &awayGoals,
		Status:          "FINISHED",
		StartTime:       time.Date(Year, 8, 10, 15, 0, 0, 0, time.UTC),
		Matchday:        &matchday,
		Stage:           "REGULAR_SEASON",
	}
}

// NewGoalScorer creates Bukayo Saka goal scorer.
func NewGoalScorer() database.GoalScorer {
	return database.GoalScorer{
		GoalScorerID:      SakaID,
		CompetitionID:     PremierLeagueID,
		StartSeasonYear:   Year,
		TeamID:            ArsenalID,
		Name:              "Bukayo Saka",
		NationalityAreaID: EnglandAreaID,
		Goals:             10,
		Assists:           5,
		GoalsFromPenalty:  0,
	}
}

func PointerToInt(value int) *int {
	return &value
}
