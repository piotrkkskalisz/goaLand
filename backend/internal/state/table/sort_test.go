package table

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func randomPermutation(clubs []*Club) []*Club {
	permutation := make([]*Club, len(clubs))
	for i, club := range clubs {
		clone := *club
		clone.Form = slices.Clone(club.Form)
		permutation[i] = &clone
	}

	rand.Shuffle(len(permutation), func(i, j int) {
		permutation[i], permutation[j] = permutation[j], permutation[i]
	})
	return permutation
}

func sortTestHelper(t *testing.T, clubs []*Club) {
	t.Helper()
	table := Table{
		Clubs: randomPermutation(clubs),
	}

	table.Sort()

	require.Equal(t, clubs, table.Clubs)
}

func TestPoints(t *testing.T) {
	clubs := []*Club{
		{
			TeamID: 1,
			Points: 8,
		}, {
			TeamID: 2,
			Points: 7,
		}, {
			TeamID: 3,
			Points: 6,
		}, {
			TeamID: 4,
			Points: 3,
		}, {
			TeamID: 5,
			Points: 2,
		},
	}
	sortTestHelper(t, clubs)
}

func TestSamePoints(t *testing.T) {
	clubs := []*Club{
		{
			TeamID:        1,
			Points:        8,
			GoalsScored:   10,
			GoalsConceded: 5,
		}, {
			TeamID:        2,
			Points:        8,
			GoalsScored:   8,
			GoalsConceded: 3,
		}, {
			TeamID:        3,
			Points:        8,
			GoalsScored:   15,
			GoalsConceded: 12,
		}, {
			TeamID:        4,
			Points:        7,
			GoalsScored:   9,
			GoalsConceded: 1,
		}, {
			TeamID:        5,
			Points:        7,
			GoalsScored:   6,
			GoalsConceded: 2,
		},
	}
	sortTestHelper(t, clubs)
}

func TestSameStatistics(t *testing.T) {
	clubs := []*Club{
		{
			TeamID:        2,
			TeamName:      "Real",
			Points:        4,
			GoalsScored:   4,
			GoalsConceded: 3,
		}, {
			TeamID:        3,
			TeamName:      "Barcelona",
			Points:        4,
			GoalsScored:   2,
			GoalsConceded: 2,
		}, {
			TeamID:        1,
			TeamName:      "Valencia",
			Points:        4,
			GoalsScored:   2,
			GoalsConceded: 2,
		},
	}
	sortTestHelper(t, clubs)
}
