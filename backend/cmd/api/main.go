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
		if err := db.ResetDatabase(); err != nil {
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
