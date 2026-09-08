package main

import (
	"backend/internal/api"
	"backend/internal/state"
	"backend/internal/sync"
	"context"
	"log"

	"backend/internal/database"
	"backend/internal/server"
)

const (
	resetDatabaseOnStart = false
	tryFetchData         = false
)

func main() {
	db, err := database.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	if resetDatabaseOnStart {
		if err := resetDatabase(db); err != nil {
			log.Fatal(err)
		}
	}

	if tryFetchData {
		if err := fetchData(db); err != nil {
			log.Fatal(err)
		}
	}

	ctx := context.Background()
	store, err := state.CreateStore(ctx, db)

	server := server.NewServer(db, store)
	if err != nil {
		log.Fatal(store)
	}

	log.Println("Starting server on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func resetDatabase(db *database.Client) error {
	if err := db.DB().Migrator().DropTable(
		&database.GoalScorer{},
		&database.Match{},
		&database.Edition{},
		&database.Team{},
		&database.Competition{},
		&database.Area{},
	); err != nil {
		return err
	}

	return db.DB().AutoMigrate(
		&database.Area{},
		&database.Competition{},
		&database.Edition{},
		&database.Team{},
		&database.Match{},
		&database.GoalScorer{},
	)
}

func fetchData(db *database.Client) error {
	api := api.NewClientFromEnv()
	ctx := context.Background()
	worker := sync.New(api, db)
	targets := []sync.SeasonTarget{{
		CompetitionCode: "PL",
		StartYear:       2026,
	}, {
		CompetitionCode: "PD",
		StartYear:       2026,
	}, {
		CompetitionCode: "BL1",
		StartYear:       2026,
	}, {
		CompetitionCode: "SA",
		StartYear:       2026,
	}, {
		CompetitionCode: "FL1",
		StartYear:       2026,
	},
	}

	return worker.InitializeData(ctx, targets)
}
