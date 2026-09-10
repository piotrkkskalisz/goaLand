package testutils

import (
	"backend/internal/database"
	"backend/internal/utils"
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
		Area:            Poland(),
	}
}

func PremierLeague() database.Competition {
	return database.Competition{
		CompetitionID:   PremierLeagueID,
		Name:            "Premier League",
		Code:            PremierLeagueCode,
		CompetitionType: "LEAGUE",
		AreaID:          EnglandAreaID,
		Area:            England(),
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
		Area:      England(),
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
		Area:      England(),
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
		Area:      England(),
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
		Area:      England(),
		Stadium:   "Anfield",
	}
}

// NewEdition creates 2025 Premier League edition.
func NewEdition() database.Edition {
	return database.Edition{
		CompetitionID: PremierLeagueID,
		Competition:   PremierLeague(),
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
		HomeTeam:        Arsenal(),
		AwayTeamID:      ManCityID,
		AwayTeam:        ManCity(),
		Edition:         NewEdition(),
		HomeGoals:       &homeGoals,
		AwayGoals:       &awayGoals,
		Status:          "FINISHED",
		StartTime:       time.Date(Year, 8, 10, 15, 0, 0, 0, time.UTC),
		Matchday:        &matchday,
		Stage:           "REGULAR_SEASON",
	}
}

// NewGoalScorer creates Bukayo Saka goal scorer.
func NewPlayer() database.Player {
	return database.Player{
		PlayerID:          SakaID,
		Name:              "Bukayo Saka",
		NationalityAreaID: EnglandAreaID,
		NationalityArea:   England(),
	}
}

func NewSeasonPlayer() database.SeasonPlayer {
	return database.SeasonPlayer{
		PlayerID:         SakaID,
		Player:           NewPlayer(),
		CompetitionID:    PremierLeagueID,
		StartSeasonYear:  Year,
		Edition:          NewEdition(),
		TeamID:           ArsenalID,
		Team:             Arsenal(),
		Goals:            utils.PointerToInt(10),
		Assists:          utils.PointerToInt(5),
		GoalsFromPenalty: utils.PointerToInt(0),
	}
}
