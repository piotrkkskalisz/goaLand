package table

import (
	"backend/internal/database"
	"backend/internal/utils"
	"cmp"
	"context"
	"slices"
)

type wrapedClub struct {
	points int
	goals  int
	club   *Club
}

//go:generate mockgen -source=table.go -destination=mocks/databse_mock.go -package=mocks
type Database interface {
	GetEditionUpcommingMatches(ctx context.Context, competitionID, startYear int) ([]database.Match, error)
	GetEditionLiveMatches(ctx context.Context, competitionID, startYear int) ([]database.Match, error)
	GetEditionResult(ctx context.Context, competitionID, startYear int) ([]database.Match, error)
}

type Table struct {
	CompetitionID int
	StartYear     int
	Clubs         []*Club
	matches       map[int]map[int]database.Match
	db            Database
}

func initClub(club *Club, team *database.Team, matches map[int]map[int]database.Match) *Club {
	if club != nil {
		return club
	}
	club = &Club{
		TeamID:   team.TeamID,
		TeamName: team.FullName,
		Form:     make([]MatchResult, 0),
	}
	matches[club.TeamID] = make(map[int]database.Match)
	return club

}
func addMatchToClub(club *Club, teamGoals, rivalGoals int, matchStatus string) *Club {

	club.addResult(teamGoals, rivalGoals)
	if slices.Contains(utils.FinishedMatchStatuses, matchStatus) {
		club.addResultToForm(teamGoals, rivalGoals)
	} else {
		club.IsLive = true

	}
	club.Points += points(teamGoals, rivalGoals)
	club.GoalsScored += teamGoals
	club.GoalsConceded += rivalGoals
	return club
}

func CreateTable(
	ctx context.Context, db Database,
	competitionID, startYear int,
) (*Table, error) {
	upcomingMatches, err := db.GetEditionUpcommingMatches(ctx, competitionID, startYear)
	if err != nil {
		return nil, err
	}

	liveMatches, err := db.GetEditionLiveMatches(ctx, competitionID, startYear)
	if err != nil {
		return nil, err
	}
	finishedMatches, err := db.GetEditionResult(ctx, competitionID, startYear)
	if err != nil {
		return nil, err
	}

	allMatches := slices.Concat(finishedMatches, liveMatches)
	clubsMap := make(map[int]*Club)
	matches := make(map[int]map[int]database.Match)

	for _, match := range upcomingMatches {
		clubsMap[match.HomeTeamID] = initClub(clubsMap[match.HomeTeamID], &match.HomeTeam, matches)
		clubsMap[match.AwayTeamID] = initClub(clubsMap[match.AwayTeamID], &match.AwayTeam, matches)
	}

	for _, match := range allMatches {
		club := initClub(clubsMap[match.HomeTeamID], &match.HomeTeam, matches)
		clubsMap[match.HomeTeamID] = addMatchToClub(
			club, *match.HomeGoals,
			*match.AwayGoals, match.Status,
		)

		club = initClub(clubsMap[match.AwayTeamID], &match.AwayTeam, matches)

		clubsMap[match.AwayTeamID] = addMatchToClub(
			club, *match.AwayGoals,
			*match.HomeGoals, match.Status,
		)

		matches[match.HomeTeamID][match.AwayTeamID] = match
	}
	var clubs []*Club
	for _, club := range clubsMap {
		clubs = append(clubs, club)
	}

	table := &Table{
		CompetitionID: competitionID,
		StartYear:     startYear,
		Clubs:         clubs,
		matches:       matches,
		db:            db,
	}
	table.Sort()
	return table, nil
}

func (t *Table) Sort() {
	if len(t.Clubs) < 2 {
		return
	}
	slices.SortFunc(t.Clubs, func(a, b *Club) int {
		if result := cmp.Compare(a.Points, b.Points); result != 0 {
			return result
		}

		return cmp.Compare(a.TeamName, b.TeamName)
	})

	indexFirstClubToAdd := 0
	var sortedClubs []*Club

	indexClub := 1
	for _, club := range t.Clubs[1:] {
		if club.Points != t.Clubs[indexFirstClubToAdd].Points {
			sortedClubs = append(
				sortedClubs,
				t.compareHeadToHead(indexFirstClubToAdd, indexClub-1)...,
			)
			indexFirstClubToAdd = indexClub
		}
		indexClub++
	}
	sortedClubs = append(
		sortedClubs,
		t.compareHeadToHead(indexFirstClubToAdd, indexClub-1)...,
	)

	slices.Reverse(sortedClubs)
	t.Clubs = sortedClubs

}

func compareTeam(a, b *Club) int {
	if result := cmp.Compare(a.goalDiffrent(), b.goalDiffrent()); result != 0 {
		return result
	}

	if result := cmp.Compare(a.GoalsScored, b.GoalsScored); result != 0 {
		return result
	}
	return cmp.Compare(b.TeamName, a.TeamName)
}

func (t *Table) sortNotFullPlayed(table []*Club) {
	slices.SortFunc(table, func(a, b *Club) int {
		return compareTeam(a, b)
	})
}

func (t *Table) isFullPlayed(clubs []*Club) bool {
	for _, club := range clubs {
		for _, otherClub := range clubs {
			if club.TeamID == otherClub.TeamID {
				continue
			}
			if _, exist := t.matches[club.TeamID][otherClub.TeamID]; !exist {
				return false
			}
		}
	}
	return true
}
func (t *Table) compareHeadToHead(indexFirstClubToAdd, indexLastClubToAdd int) []*Club {
	clubs := t.Clubs[indexFirstClubToAdd : indexLastClubToAdd+1]
	if indexFirstClubToAdd == indexLastClubToAdd {
		return []*Club{t.Clubs[indexFirstClubToAdd]} //add only 1 team with this number points
	}

	if !t.isFullPlayed(clubs) {
		t.sortNotFullPlayed(clubs)
		return clubs
	}

	var minTable []*wrapedClub
	for _, club := range clubs {
		minTable = append(minTable, &wrapedClub{
			club: club,
		})
	}
	//TODO: srotuj mecze
	for _, club := range minTable {
		for _, otherClub := range minTable {
			if club.club.TeamID == otherClub.club.TeamID {
				continue
			}
			match := t.matches[club.club.TeamID][otherClub.club.TeamID]
			club.points += points(*match.HomeGoals, *match.AwayGoals)
			club.goals += *match.HomeGoals - *match.AwayGoals

			otherClub.points += points(*match.AwayGoals, *match.HomeGoals)
			otherClub.goals += *match.AwayGoals - *match.HomeGoals
		}
	}

	slices.SortFunc(minTable, func(a, b *wrapedClub) int {
		if result := cmp.Compare(a.points, b.points); result != 0 {
			return result
		}

		if result := cmp.Compare(a.goals, b.goals); result != 0 {
			return result
		}

		return compareTeam(a.club, b.club)
	})

	var sortedClubs []*Club
	for _, club := range minTable {
		sortedClubs = append(sortedClubs, club.club)
	}

	return sortedClubs
}
func points(teamGoals, rivalGoals int) int {
	if teamGoals > rivalGoals {
		return 3
	} else if teamGoals == rivalGoals {
		return 1
	} else {
		return 0
	}
}
