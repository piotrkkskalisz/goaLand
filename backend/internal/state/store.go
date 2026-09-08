package state

import (
	"backend/internal/database"
	"backend/internal/state/table"
	"context"
)

type tableKey struct {
	competitionID int
	startYear     int
}

type Store struct {
	tables map[tableKey]*table.Table
	db     *database.Client
}

func CreateStore(ctx context.Context, db *database.Client) (*Store, error) {
	var editions []database.Edition

	db.List(ctx, &editions, database.Filter{})
	tables := make(map[tableKey]*table.Table)

	for _, edition := range editions {
		table, err := table.CreateTable(ctx, db, edition.CompetitionID, edition.StartYear)
		if err != nil {
			return nil, err
		}
		tables[GetKey(table)] = table
	}
	return &Store{
		tables: tables,
		db:     db,
	}, nil
}

func GetKey(table *table.Table) tableKey {
	return tableKey{
		competitionID: table.CompetitionID,
		startYear:     table.StartYear,
	}
}

func (s *Store) GetTable(competitionID int, startYear int) (*table.Table, bool) {
	key := tableKey{
		competitionID: competitionID,
		startYear:     startYear,
	}

	table, exists := s.tables[key]
	if !exists {
		return nil, false
	}
	return table, true
}
