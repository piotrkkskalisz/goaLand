package testutils

import (
	"backend/internal/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func Database_init(t *testing.T, client *database.Client) {
	t.Helper()
	require.NoError(t, client.DB().Migrator().DropTable(
		&database.GoalScorer{},
		&database.Match{},
		&database.Edition{},
		&database.Team{},
		&database.Competition{},
		&database.Area{},
	))

	require.NoError(t, client.DB().AutoMigrate(
		&database.Area{},
		&database.Competition{},
		&database.Edition{},
		&database.Team{},
		&database.Match{},
		&database.GoalScorer{},
	))
}
