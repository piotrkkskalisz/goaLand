package main

import (
	"backend/internal/api"
	"backend/internal/sync"
	"context"
	"log"

	"backend/internal/database"
	"backend/internal/server"
)

const (
	resetDatabaseOnStart = false
	tryFetchData         = true
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

	server := server.NewServer(db)

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
	target := sync.SeasonTarget{
		CompetitionCode: "PL",
		StartYear:       2026,
	}
	return worker.InitializeData(ctx, sync.SeasonTargets{target})
}
