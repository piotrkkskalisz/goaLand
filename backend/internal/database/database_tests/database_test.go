//go:build integration

package database_tests

import (
	db "backend/internal/database"
	"backend/internal/testutils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClientFromEnv(t *testing.T) {
	config, err := db.NewConfigFromEnv()
	require.NoError(t, err)

	client, err := db.NewClient(config)
	require.NoError(t, err)
	require.NotNil(t, client)

	sqlDB, err := client.DB().DB()
	require.NoError(t, err)

	require.NoError(t, sqlDB.Ping())
}

func saveAndAssertAreas(t *testing.T, client *db.Client, areas []db.Area) {
	t.Helper()

	require.NoError(t, client.SaveAreas(areas))

	var saved []db.Area
	require.NoError(t, client.DB().Find(&saved).Error)

	require.Len(t, saved, len(areas))

	require.ElementsMatch(t, areas, saved)
}

func TestSaveAreas(t *testing.T) {
	client, err := db.NewClientFromEnv()
	require.NoError(t, err)

	require.NoError(t, client.DB().Migrator().DropTable(&db.Area{}))
	require.NoError(t, client.DB().AutoMigrate(&db.Area{}))

	areas := []db.Area{testutils.England(), testutils.Poland()}

	saveAndAssertAreas(t, client, areas)

	areas[1].Name = "Germany"
	areas[1].Code = "GER"
	saveAndAssertAreas(t, client, areas)

}
