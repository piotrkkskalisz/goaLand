package database

import "time"

type Area struct {
	AreaID    int `gorm:"primaryKey"`
	Name      string
	Code      string
	IsCountry bool

	Competitions []Competition
	Teams        []Team
	GoalScorers  []GoalScorer `gorm:"foreignKey:NationalityAreaID"`
}

type Competition struct {
	CompetitionID int `gorm:"primaryKey"`

	Name            string
	Code            string
	CompetitionType string

	AreaID int
	Area   Area

	Editions []Edition
}

type Team struct {
	TeamID int `gorm:"primaryKey"`

	FullName  string
	ShortName string
	Code      string
	Colors    string

	Stadium string
	AreaID  int
	Area    Area

	HomeMatches []Match `gorm:"foreignKey:HomeTeamID"`
	AwayMatches []Match `gorm:"foreignKey:AwayTeamID"`

	GoalScorers []GoalScorer
}

type Edition struct {
	CompetitionID int `gorm:"primaryKey"`
	Competition   Competition

	StartYear int `gorm:"primaryKey"`
	Status    string

	Matches     []Match      `gorm:"foreignKey:CompetitionID,StartSeasonYear;references:CompetitionID,StartYear"`
	GoalScorers []GoalScorer `gorm:"foreignKey:CompetitionID,StartSeasonYear;references:CompetitionID,StartYear"`
}

type Match struct {
	MatchID int `gorm:"primaryKey"`

	HomeTeamID int
	HomeTeam   Team `gorm:"foreignKey:HomeTeamID"`

	AwayTeamID int
	AwayTeam   Team `gorm:"foreignKey:AwayTeamID"`

	CompetitionID   int
	StartSeasonYear int

	Edition Edition `gorm:"foreignKey:CompetitionID,StartSeasonYear;references:CompetitionID,StartYear"`

	HomeGoals *int
	AwayGoals *int

	HalfTimeHomeGoals *int
	HalfTimeAwayGoals *int

	Status    string
	StartTime time.Time
}

type GoalScorer struct {
	GoalScorerID int `gorm:"primaryKey"`

	TeamID int
	Team   Team

	CompetitionID   int `gorm:"primaryKey"`
	StartSeasonYear int `gorm:"primaryKey"`

	Edition Edition `gorm:"foreignKey:CompetitionID,StartSeasonYear;references:CompetitionID,StartYear"`

	Name string

	NationalityAreaID int
	NationalityArea   Area `gorm:"foreignKey:NationalityAreaID"`

	Goals            int
	Assists          int
	GoalsFromPenalty int
}
