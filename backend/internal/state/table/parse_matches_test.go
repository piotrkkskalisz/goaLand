package table

import (
	"backend/internal/database"
	"backend/internal/state/table/mocks"
	"backend/internal/testutils"
	"backend/internal/utils"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func tableTestHelper(t *testing.T, finishedMatches, liveMatches, upcomingMatches []database.Match, expectedClubs []*Club) {
	t.Helper()

	reversedFinishedMatches := make([]database.Match, len(finishedMatches))
	// 2. Kopiujemy elementy z oryginalnego slice'a
	copy(reversedFinishedMatches, finishedMatches)
	// 3. Odwracamy kopię w miejscu
	slices.Reverse(reversedFinishedMatches)

	ctrl := gomock.NewController(t)
	dbMock := mocks.NewMockDatabase(ctrl)
	dbMock.EXPECT().GetEditionResult(t.Context(), 1, 2025).Return(reversedFinishedMatches, nil)
	dbMock.EXPECT().GetEditionLiveMatches(t.Context(), 1, 2025).Return(liveMatches, nil)
	dbMock.EXPECT().GetEditionUpcommingMatches(t.Context(), 1, 2025).Return(upcomingMatches, nil)

	table, err := CreateTable(t.Context(), dbMock, 1, 2025)

	require.NoError(t, err)
	require.Equal(t, expectedClubs, table.Clubs)
}
func TestFullMatch(t *testing.T) {
	arsenal := testutils.Arsenal()
	manCity := testutils.ManCity()
	chelsea := testutils.Chelsea()

	matches := []database.Match{
		{
			MatchID:    1,
			HomeTeamID: arsenal.TeamID,
			HomeTeam:   arsenal,
			AwayTeamID: manCity.TeamID,
			AwayTeam:   manCity,
			HomeGoals:  utils.PointerToInt(2),
			AwayGoals:  utils.PointerToInt(1),
			Status:     "FINISHED",
		}, {
			MatchID:    2,
			HomeTeamID: manCity.TeamID,
			HomeTeam:   manCity,
			AwayTeamID: arsenal.TeamID,
			AwayTeam:   arsenal,
			HomeGoals:  utils.PointerToInt(2),
			AwayGoals:  utils.PointerToInt(0),
			Status:     "FINISHED",
		}, {
			MatchID:    3,
			HomeTeamID: arsenal.TeamID,
			HomeTeam:   arsenal,
			AwayTeamID: chelsea.TeamID,
			AwayTeam:   chelsea,
			HomeGoals:  utils.PointerToInt(5),
			AwayGoals:  utils.PointerToInt(1),
			Status:     "FINISHED",
		}, {
			MatchID:    4,
			HomeTeamID: chelsea.TeamID,
			HomeTeam:   chelsea,
			AwayTeamID: arsenal.TeamID,
			AwayTeam:   arsenal,
			HomeGoals:  utils.PointerToInt(2),
			AwayGoals:  utils.PointerToInt(2),
			Status:     "FINISHED",
		}, {
			MatchID:    5,
			HomeTeamID: manCity.TeamID,
			HomeTeam:   manCity,
			AwayTeamID: chelsea.TeamID,
			AwayTeam:   chelsea,
			HomeGoals:  utils.PointerToInt(3),
			AwayGoals:  utils.PointerToInt(1),
			Status:     "FINISHED",
		}, {
			MatchID:    6,
			HomeTeamID: chelsea.TeamID,
			HomeTeam:   chelsea,
			AwayTeamID: manCity.TeamID,
			AwayTeam:   manCity,
			HomeGoals:  utils.PointerToInt(2),
			AwayGoals:  utils.PointerToInt(2),
			Status:     "FINISHED",
		},
	}
	// Arsenal and City have 7 points and a +3 goal difference. Arsenal has
	// scored more goals, but City ranks first thanks to the 3:2 head-to-head score.
	expectedClubs := []*Club{
		{
			TeamID:        manCity.TeamID,
			TeamName:      manCity.FullName,
			Points:        7,
			Wins:          2,
			Draws:         1,
			Losses:        1,
			GoalsScored:   8,
			GoalsConceded: 5,
			Form:          []MatchResult{Loss, Win, Win, Draw},
		}, {
			TeamID:        arsenal.TeamID,
			TeamName:      arsenal.FullName,
			Points:        7,
			Wins:          2,
			Draws:         1,
			Losses:        1,
			GoalsScored:   9,
			GoalsConceded: 6,
			Form:          []MatchResult{Win, Loss, Win, Draw},
		}, {
			TeamID:        chelsea.TeamID,
			TeamName:      chelsea.FullName,
			Points:        2,
			Wins:          0,
			Draws:         2,
			Losses:        2,
			GoalsScored:   6,
			GoalsConceded: 12,
			Form:          []MatchResult{Loss, Draw, Loss, Draw},
		},
	}

	tableTestHelper(t, matches, nil, nil, expectedClubs)
}

func TestPartMatch(t *testing.T) {
	arsenal := testutils.Arsenal()
	manCity := testutils.ManCity()
	chelsea := testutils.Chelsea()

	finishedMatches := []database.Match{
		{
			MatchID:    1,
			HomeTeamID: arsenal.TeamID,
			HomeTeam:   arsenal,
			AwayTeamID: manCity.TeamID,
			AwayTeam:   manCity,
			HomeGoals:  utils.PointerToInt(1),
			AwayGoals:  utils.PointerToInt(0),
			Status:     "FINISHED",
		},
	}
	liveMatches := []database.Match{
		{
			MatchID:    2,
			HomeTeamID: manCity.TeamID,
			HomeTeam:   manCity,
			AwayTeamID: chelsea.TeamID,
			AwayTeam:   chelsea,
			HomeGoals:  utils.PointerToInt(4),
			AwayGoals:  utils.PointerToInt(0),
			Status:     "IN_PLAY",
		},
	}
	upcomingMatches := []database.Match{
		{
			MatchID:    4,
			HomeTeamID: manCity.TeamID,
			HomeTeam:   manCity,
			AwayTeamID: arsenal.TeamID,
			AwayTeam:   arsenal,
			Status:     "SCHEDULED",
		},
		{
			MatchID:    3,
			HomeTeamID: chelsea.TeamID,
			HomeTeam:   chelsea,
			AwayTeamID: arsenal.TeamID,
			AwayTeam:   arsenal,
			Status:     "SCHEDULED",
		},
	}
	expectedClubs := []*Club{
		{
			TeamID:        manCity.TeamID,
			TeamName:      manCity.FullName,
			Points:        3,
			Wins:          1,
			Losses:        1,
			GoalsScored:   4,
			GoalsConceded: 1,
			Form:          []MatchResult{Loss},
			IsLive:        true,
		},
		{
			TeamID:      arsenal.TeamID,
			TeamName:    arsenal.FullName,
			Points:      3,
			Wins:        1,
			GoalsScored: 1,
			Form:        []MatchResult{Win},
		},
		{
			TeamID:        chelsea.TeamID,
			TeamName:      chelsea.FullName,
			Points:        0,
			Losses:        1,
			GoalsConceded: 4,
			Form:          []MatchResult{},
			IsLive:        true,
		},
	}
	tableTestHelper(t, finishedMatches, liveMatches, upcomingMatches, expectedClubs)

}
