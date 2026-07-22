//go:build integration

package sync

import (
	"backend/internal/database"
	"backend/internal/testutils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitializeData(t *testing.T) {
	dbClient, err := database.NewClientFromEnv()
	require.NoError(t, err)

	testutils.Database_init(t, dbClient)

	s, err := NewFromEnv()
	require.NoError(t, err)

	previosYear := testutils.Year - 1
	err = s.InitializeData(SeasonTargets{
		{
			CompetitionCode: testutils.PremierLeagueCode,
			StartYear:       testutils.Year,
		}, {
			CompetitionCode: testutils.PremierLeagueCode,
			StartYear:       previosYear,
		},
	})
	require.NoError(t, err)
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
	require.Len(t, editions, 1)

	var teams []database.Team
	require.NoError(t, dbClient.DB().Find(&teams).Error)
	require.NotEmpty(t, teams)

	var matches []database.Match
	require.NoError(t, dbClient.DB().Find(&matches).Error)
	require.NotEmpty(t, matches)

	var scorers []database.GoalScorer
	require.NoError(t, dbClient.DB().Find(&scorers).Error)
	require.NotEmpty(t, scorers)

	require.Equal(t, "Premier League", competitions[0].Name)
	require.Equal(t, testutils.Year, editions[0].StartYear)

	require.Contains(t, []string{"England", "Poland"}, areas[0].Name)
}
