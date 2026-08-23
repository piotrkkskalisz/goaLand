package transform

import (
	"backend/internal/database"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	year    = 2026
	month   = time.August
	today   = 10
	nowHour = 19
)

var now = time.Date(year, month, today, nowHour, 0, 0, 0, time.UTC)

func testMatch(id int, day int, hour int, status string) database.Match {
	return database.Match{
		MatchID:   id,
		StartTime: time.Date(year, month, day, hour, 0, 0, 0, time.UTC),
		Status:    status,
	}
}

func matchIDs(matches []database.Match) []int {
	ids := make([]int, 0, len(matches))
	for _, match := range matches {
		ids = append(ids, match.MatchID)
	}
	return ids
}

func TestAllTodayMatches(t *testing.T) {
	matches := []database.Match{
		testMatch(1, 10, 10, "FINISHED"),
		testMatch(2, 10, 15, "TIMED"),
		testMatch(3, 10, 18, "TIMED"),
		testMatch(4, 11, 15, "TIMED"),
	}
	featured := chooseFeatured(matches, now)
	require.Equal(t, []int{1, 2, 3}, matchIDs(featured))
}

func TestLastFinishedMatch(t *testing.T) {
	matches := []database.Match{
		testMatch(1, 8, 18, "FINISHED"),
		testMatch(2, 9, 18, "FINISHED"),
		testMatch(3, 10, 21, "TIMED"),
	}

	featured := chooseFeatured(matches, now)
	require.Equal(t, []int{2, 3}, matchIDs(featured))
}

func TestAddsOneNextMatch(t *testing.T) {
	matches := []database.Match{
		testMatch(1, 10, 12, "FINISHED"),
		testMatch(2, 10, 15, "FINISHED"),
		testMatch(3, 10, 18, "FINISHED"),
		testMatch(4, 11, 15, "TIMED"),
		testMatch(5, 11, 18, "TIMED"),
	}

	featured := chooseFeatured(matches, now)

	require.Equal(t, []int{1, 2, 3, 4}, matchIDs(featured))
}

func TestAddsTwoNextMatches(t *testing.T) {
	matches := []database.Match{
		testMatch(4, 9, 17, "FINISHED"),
		testMatch(1, 9, 18, "FINISHED"),
		testMatch(2, 11, 15, "TIMED"),
		testMatch(3, 11, 18, "TIMED"),
		testMatch(5, 11, 19, "TIMED"),
	}

	featured := chooseFeatured(matches, now)

	require.Equal(t, []int{1, 2, 3}, matchIDs(featured))
}

func TestLiveMatchFromPreviousDay(t *testing.T) {
	now := time.Date(year, month, today, 1, 0, 0, 0, time.UTC)
	matches := []database.Match{
		testMatch(1, 8, 18, "FINISHED"),
		testMatch(2, 9, 23, "IN_PLAY"),
		testMatch(3, 11, 15, "TIMED"),
	}

	featured := chooseFeatured(matches, now)

	require.Equal(t, []int{1, 2, 3}, matchIDs(featured))
}

func TestEmptyMatches(t *testing.T) {
	featured := chooseFeatured(nil, now)

	require.NotNil(t, featured)
	require.Empty(t, featured)
}
