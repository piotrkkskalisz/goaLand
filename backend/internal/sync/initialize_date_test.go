//go:build integration

package sync

import (
	"backend/internal/database"
	"backend/internal/testutils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitializeData(t *testing.T) {
	ctx := t.Context()

	dbClient, err := database.NewClientFromEnv()
	require.NoError(t, err)

	require.NoError(t, dbClient.DeleteDatabase())
	require.NoError(t, dbClient.CreateDatabase())

	s, err := NewFromEnv()
	require.NoError(t, err)

	previousYear := testutils.Year - 1
	err = s.InitializeData(ctx, SeasonTargets{
		{
			CompetitionCode: testutils.PremierLeagueCode,
			StartYear:       testutils.Year,
		}, {
			CompetitionCode: testutils.PremierLeagueCode,
			StartYear:       previousYear,
		},
	})
	require.NoError(t, err)

	verifyDatas(t, dbClient)
}

func verifyDatas(t *testing.T, dbClient *database.Client) {
	t.Helper()

	var areas []database.Area
	require.NoError(t, dbClient.DB().Find(&areas).Error)
	require.NotEmpty(t, areas)

	var competitions []database.Competition
	require.NoError(t, dbClient.DB().Find(&competitions).Error)
	require.Len(t, competitions, 1)

	var editions []database.Edition
	require.NoError(t, dbClient.DB().Find(&editions).Error)
	require.Len(t, editions, 2)

	var teams []database.Team
	require.NoError(t, dbClient.DB().Find(&teams).Error)
	require.NotEmpty(t, teams)

	var matches []database.Match
	require.NoError(t, dbClient.DB().Find(&matches).Error)
	require.NotEmpty(t, matches)

	var players []database.Player
	require.NoError(t, dbClient.DB().Find(&players).Error)
	require.NotEmpty(t, players)

	var seasonPlayers []database.SeasonPlayer
	require.NoError(t, dbClient.DB().Find(&seasonPlayers).Error)
	require.NotEmpty(t, seasonPlayers)

	require.Equal(t, "Premier League", competitions[0].Name)
	require.ElementsMatch(t,
		[]int{testutils.Year, testutils.Year - 1},
		[]int{editions[0].StartYear, editions[1].StartYear},
	)

	areaNames := make([]string, 0, len(areas))
	for _, area := range areas {
		areaNames = append(areaNames, area.Name)
	}
	require.Contains(t, areaNames, testutils.EnglandName)
}
