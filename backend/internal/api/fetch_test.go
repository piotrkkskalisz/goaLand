//go:build integration

package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

const PremierLeagueCode = "PL"

func nilOrString(t *testing.T, ptr *string) string {
	t.Helper()
	value := "<nil>"
	if ptr != nil {
		value = *ptr
		require.NotEmpty(t, value)
	}
	return value
}

func nilOrInt(ptr *int) string {
	if ptr != nil {
		return strconv.Itoa(*ptr)
	}
	return "<nil>"
}

func TestConnection(t *testing.T) {
	client := NewClientFromEnv()

	_, err := client.FetchAreas()

	require.NoError(t, err)
}

func TestFetchAreas(t *testing.T) {
	client := NewClientFromEnv()

	areas, err := client.FetchAreas()

	require.NoError(t, err)
	require.NotEmpty(t, areas)

	for _, area := range areas {
		parent := nilOrString(t, area.ParentArea)

		t.Logf("Area: %+v, ParentArea: %s", area, parent)

		require.NotZero(t, area.ID)
		require.NotEmpty(t, area.Name)
		require.NotEmpty(t, area.CountryCode)
	}
}

func TestFetchEdition(t *testing.T) {
	client := NewClientFromEnv()

	edition, err := client.FetchEdition(PremierLeagueCode, 2025)

	require.NoError(t, err)

	t.Logf("Edition: %+v", edition)

	require.NotZero(t, edition.ID)
	require.NotEmpty(t, edition.StartDate)
	require.NotEmpty(t, edition.EndDate)
}

func TestFetchTeams(t *testing.T) {
	client := NewClientFromEnv()

	teams, err := client.FetchTeams(PremierLeagueCode, 2025)

	require.NoError(t, err)
	require.NotEmpty(t, teams)

	for _, team := range teams {
		t.Logf("Team: %+v", team)

		require.NotZero(t, team.ID)
		require.NotZero(t, team.Area.ID)

		require.NotEmpty(t, team.Name)
		require.NotEmpty(t, team.ShortName)
		require.NotEmpty(t, team.TLA)
		require.NotEmpty(t, team.ClubColors)
		require.NotEmpty(t, team.Venue)
	}
}

func TestFetchMatches(t *testing.T) {
	client := NewClientFromEnv()

	worldCupCode := "WC"
	matches, err := client.FetchMatches(worldCupCode, 2026)

	require.NoError(t, err)
	require.NotEmpty(t, matches)

	for _, match := range matches {
		fullTimeHome := nilOrInt(match.Score.FullTime.Home)
		fullTimeAway := nilOrInt(match.Score.FullTime.Away)
		halfTimeHome := nilOrInt(match.Score.HalfTime.Home)
		halfTimeAway := nilOrInt(match.Score.HalfTime.Away)

		t.Logf(
			"Match: %+v, FullTime: %s:%s, HalfTime: %s:%s",
			match,
			fullTimeHome,
			fullTimeAway,
			halfTimeHome,
			halfTimeAway,
		)

		require.NotZero(t, match.ID)

		require.NotEmpty(t, match.UtcDate)
		require.NotEmpty(t, match.Status)

		require.NotZero(t, match.HomeTeam.ID)
		require.NotZero(t, match.AwayTeam.ID)

		if match.Score.FullTime.Home != nil {
			require.GreaterOrEqual(t, *match.Score.FullTime.Home, 0)
		}
		if match.Score.FullTime.Away != nil {
			require.GreaterOrEqual(t, *match.Score.FullTime.Away, 0)
		}
		if match.Score.HalfTime.Home != nil {
			require.GreaterOrEqual(t, *match.Score.HalfTime.Home, 0)
		}
		if match.Score.HalfTime.Away != nil {
			require.GreaterOrEqual(t, *match.Score.HalfTime.Away, 0)
		}
	}
}

func TestFetchGoalScorers(t *testing.T) {
	client := NewClientFromEnv()

	scorers, err := client.FetchGoalScorers(PremierLeagueCode, 2025, 10)

	require.NoError(t, err)
	require.NotEmpty(t, scorers)

	for _, scorer := range scorers {
		assists := nilOrInt(scorer.Assists)
		penalties := nilOrInt(scorer.Penalties)

		t.Logf(
			"GoalScorer: %+v, Assists: %s, Penalties: %s",
			scorer,
			assists,
			penalties,
		)

		require.NotZero(t, scorer.Player.ID)
		require.NotZero(t, scorer.Team.ID)

		require.GreaterOrEqual(t, scorer.Goals, 0)

		require.NotEmpty(t, scorer.Player.Nationality)
		require.NotEmpty(t, scorer.Player.Name)

		if scorer.Assists != nil {
			require.GreaterOrEqual(t, *scorer.Assists, 0)
		}

		if scorer.Penalties != nil {
			require.GreaterOrEqual(t, *scorer.Penalties, 0)
		}
	}
}
