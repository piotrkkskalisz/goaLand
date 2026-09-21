package main

import (
	"backend/internal/api"
	"backend/internal/state"
	"backend/internal/sync"
	"context"
	"log"
	"time"

	"backend/internal/database"
	"backend/internal/server"
)

const (
	tryFetchData = false

	tryTrackData = true
)

var trackedCompetitionCodes = []string{"PL", "PD", "BL1", "SA", "FL1"}

func main() {
	ctx := context.Background()

	db, err := database.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	store := state.NewStore(ctx, db)

	apiClient := api.NewClientFromEnv()

	worker := sync.New(apiClient, db, store)
	if tryFetchData {
		if err := db.ResetDatabase(); err != nil {
			log.Fatal(err)
		}
		if err := fetchData(ctx, worker); err != nil {
			log.Fatal(err)
		}
	} else {
		worker.InitAreasFromDB(ctx)
	}

	if err = store.InitTables(ctx); err != nil {
		log.Fatal(err)
	}

	if tryTrackData {
		if err := worker.TrackCompetitions(trackedCompetitionCodes); err != nil {
			log.Fatal(err)
		}
	}

	server := server.NewServer(db, store)

	log.Println("Starting server on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func fetchData(ctx context.Context, worker *sync.Sync) error {
	now := time.Now()

	oldSeasonTargets := []sync.SeasonTarget{{
		CompetitionCode: "PL",
		StartYear:       2025,
	}}

	err := worker.InitializeData(ctx, oldSeasonTargets)
	if err != nil {
		return err
	}

	for _, code := range trackedCompetitionCodes {
		worker.AddSeason(ctx, now, sync.SeasonTarget{
			CompetitionCode: code,
			StartYear:       sync.CurrentSeasonStartYear(now),
		})
	}
	return nil
}
