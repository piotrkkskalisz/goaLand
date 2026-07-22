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

	require.NoError(t, client.SaveAreas([]db.Area{area}))
	require.NoError(t, client.SaveCompetitions([]db.Competition{competition}))
	require.NoError(t, client.SaveTeams([]db.Team{team, secondTeam}))
	require.NoError(t, client.SaveEdition(edition))
	require.NoError(t, client.SaveMatches([]db.Match{match}))
	require.NoError(t, client.SaveGoalScorers([]db.GoalScorer{goalScorer}))

	var loadedEdition db.Edition
	require.NoError(t, client.DB().
		Preload("Competition").
		Preload("Matches").
		Preload("GoalScorers").
		First(&loadedEdition).Error)

	require.Equal(t, competition.Name, loadedEdition.Competition.Name)
	require.Len(t, loadedEdition.Matches, 1)
	require.Len(t, loadedEdition.GoalScorers, 1)
}
