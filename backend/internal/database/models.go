package database

import "time"

type Area struct {
	AreaID    int `gorm:"primaryKey"`
	Name      string `gorm:"not null;unique"`
	Code      string `gorm:"not null;unique"`
	IsCountry bool   `gorm:"not null"`

	Competitions []Competition
	Teams        []Team
	GoalScorers  []GoalScorer `gorm:"foreignKey:NationalityAreaID"`
}

type Competition struct {
	CompetitionID int `gorm:"primaryKey"`

	Name            string `gorm:"not null;unique"`
	Code            string `gorm:"not null;unique"`
	CompetitionType string `gorm:"not null"`

	AreaID int `gorm:"not null"`
	Area   Area

	Editions []Edition
}

type Team struct {
	TeamID int `gorm:"primaryKey"`

	FullName  string `gorm:"not null;unique"`
	ShortName string `gorm:"not null"`
	Code      string `gorm:"not null;unique"`
	Colors    string

	Stadium string `gorm:"not null"`
	AreaID  int    `gorm:"not null"`
	Area    Area

	HomeMatches []Match `gorm:"foreignKey:HomeTeamID"`
	AwayMatches []Match `gorm:"foreignKey:AwayTeamID"`

	GoalScorers []GoalScorer
}

type Edition struct {
	CompetitionID int `gorm:"primaryKey"`
	Competition   Competition

	StartYear int `gorm:"primaryKey"`
	Status    string `gorm:"not null"`

	Matches     []Match      `gorm:"foreignKey:CompetitionID,StartSeasonYear;references:CompetitionID,StartYear"`
	GoalScorers []GoalScorer `gorm:"foreignKey:CompetitionID,StartSeasonYear;references:CompetitionID,StartYear"`
}

type Match struct {
	MatchID int `gorm:"primaryKey"`

	HomeTeamID int  `gorm:"not null"`
	HomeTeam   Team `gorm:"foreignKey:HomeTeamID"`

	AwayTeamID int  `gorm:"not null"`
	AwayTeam   Team `gorm:"foreignKey:AwayTeamID"`

	CompetitionID   int `gorm:"not null"`
	StartSeasonYear int `gorm:"not null"`

	Edition Edition `gorm:"foreignKey:CompetitionID,StartSeasonYear;references:CompetitionID,StartYear"`

	HomeGoals *int
	AwayGoals *int

	HalfTimeHomeGoals *int
	HalfTimeAwayGoals *int

	Status    string    `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`

	Matchday *int
	Stage    string `gorm:"not null"`
}

type GoalScorer struct {
	GoalScorerID int `gorm:"primaryKey"`

	TeamID int `gorm:"not null"`
	Team   Team

	CompetitionID   int `gorm:"primaryKey"`
	StartSeasonYear int `gorm:"primaryKey"`

	Edition Edition `gorm:"foreignKey:CompetitionID,StartSeasonYear;references:CompetitionID,StartYear"`

	Name string `gorm:"not null"`

	NationalityAreaID int `gorm:"not null"`
	NationalityArea   Area `gorm:"foreignKey:NationalityAreaID"`

	Goals            int `gorm:"not null"`
	Assists          int `gorm:"not null"`
	GoalsFromPenalty int `gorm:"not null"`
}
