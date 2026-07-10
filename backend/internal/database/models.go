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

type Stadium struct {
	StadiumID int `gorm:"primaryKey"`
	Name      string

	Teams   []Team
	Matches []Match
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

type Edition struct {
	EditionID int `gorm:"primaryKey"`

	CompetitionID int
	Competition   Competition

	StartYear int
	Status    string

	Matches     []Match
	GoalScorers []GoalScorer
}

type Team struct {
	TeamID int `gorm:"primaryKey"`

	FullName  string
	ShortName string
	Code      string
	Colors    string

	StadiumID *int
	Stadium   *Stadium

	AreaID int
	Area   Area

	HomeMatches []Match `gorm:"foreignKey:HomeTeamID"`
	AwayMatches []Match `gorm:"foreignKey:AwayTeamID"`

	GoalScorers []GoalScorer
}

type Match struct {
	MatchID int `gorm:"primaryKey"`

	HomeTeamID int
	HomeTeam   Team `gorm:"foreignKey:HomeTeamID"`

	AwayTeamID int
	AwayTeam   Team `gorm:"foreignKey:AwayTeamID"`

	StadiumID *int
	Stadium   *Stadium

	EditionID int
	Edition   Edition

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

	EditionID int
	Edition   Edition

	Name string

	NationalityAreaID int
	NationalityArea   Area `gorm:"foreignKey:NationalityAreaID"`

	Position string

	Goals            int
	Assists          int
	GoalsFromPenalty int
}
