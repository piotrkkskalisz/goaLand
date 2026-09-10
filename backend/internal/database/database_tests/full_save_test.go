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

	require.NoError(t, client.DeleteDatabase())
	require.NoError(t, client.CreateDatabase())

	area := testutils.England()
	competition := testutils.NewCompetition()
	edition := testutils.NewEdition()

	team := testutils.Arsenal()
	secondTeam := testutils.ManCity()

	match := testutils.NewMatch()
	player := testutils.NewPlayer()
	seasonPlayer := testutils.NewSeasonPlayer()

	match.StartTime = time.Now()

	ctx := t.Context()

	require.NoError(t, client.Save(ctx, []db.Area{area}))
	require.NoError(t, client.Save(ctx, []db.Competition{competition}))
	require.NoError(t, client.Save(ctx, []db.Team{team, secondTeam}))
	require.NoError(t, client.Save(ctx, &edition))
	require.NoError(t, client.Save(ctx, []db.Match{match}))
	require.NoError(t, client.Save(ctx, []db.Player{player}))
	require.NoError(t, client.Save(ctx, []db.SeasonPlayer{seasonPlayer}))

	var loadedEdition db.Edition
	require.NoError(t, client.DB().
		Preload("Competition").
		Preload("Matches").
		Preload("SeasonPlayers.Player").
		First(&loadedEdition).Error)

	require.Equal(t, competition.Name, loadedEdition.Competition.Name)
	require.Len(t, loadedEdition.Matches, 1)
	require.Equal(t, match.Matchday, loadedEdition.Matches[0].Matchday)
	require.Equal(t, match.Stage, loadedEdition.Matches[0].Stage)
	require.Len(t, loadedEdition.SeasonPlayers, 1)
	require.Equal(t, player.PlayerID, loadedEdition.SeasonPlayers[0].PlayerID)
	require.Equal(t, player.Name, loadedEdition.SeasonPlayers[0].Player.Name)
	require.Equal(t, seasonPlayer.Goals, loadedEdition.SeasonPlayers[0].Goals)
}
