//go:build integration

package database_tests

import (
	"testing"
	"time"

	db "backend/internal/database"
	"backend/internal/testutils"

	"github.com/stretchr/testify/require"
)

func TestDatabaseIntegrity(t *testing.T) {
	client, err := db.NewClientFromEnv()
	require.NoError(t, err)

	testutils.Arsenal()
	testutils.Database_init(t, client)

	area := testutils.England()
	competition := testutils.NewCompetition()
	edition := testutils.NewEdition()

	team := testutils.Arsenal()
	secondTeam := testutils.ManCity()

	match := testutils.NewMatch()
	goalScorer := testutils.NewGoalScorer()

	match.StartTime = time.Now()

	ctx := t.Context()

	require.NoError(t, client.Save(ctx, []db.Area{area}))
	require.NoError(t, client.Save(ctx, []db.Competition{competition}))
	require.NoError(t, client.Save(ctx, []db.Team{team, secondTeam}))
	require.NoError(t, client.Save(ctx, edition))
	require.NoError(t, client.Save(ctx, []db.Match{match}))
	require.NoError(t, client.Save(ctx, []db.GoalScorer{goalScorer}))

	var loadedEdition db.Edition
	require.NoError(t, client.DB().
		Preload("Competition").
		Preload("Matches").
		Preload("GoalScorers").
		First(&loadedEdition).Error)

	require.Equal(t, competition.Name, loadedEdition.Competition.Name)
	require.Len(t, loadedEdition.Matches, 1)
	require.Equal(t, match.Matchday, loadedEdition.Matches[0].Matchday)
	require.Equal(t, match.Stage, loadedEdition.Matches[0].Stage)
	require.Len(t, loadedEdition.GoalScorers, 1)
}
