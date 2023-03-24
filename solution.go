package main

import (
	"context"
	"errors"
	"events-dedup/internal/database"
	"os"
)

var DuplicateEventError = errors.New("duplicate event")

// deduplicateEvent - checks if the event was processed before.
// Solution should be persisted under application restarts.
func deduplicateEvent(event *Event) (*Event, error) {
	// panic("not implemented")
	return solution(event)
}

var (
	db  database.Repository = database.NewRepository(ctx)
	ctx                     = context.Background()
)

func init() {
	initBloom()
	initNaive()
	db.RunMigrations()
	events, err := db.ListEvents(ctx)
	if err != nil {
		panic(err)
	}
	for _, event := range events {
		switch os.Getenv("SOLUTION") {
		case "bloom":
			bloomFilter.Add([]byte(event.Hash))
		case "naive":
			fallthrough
		default:
			processedQueue[event.Hash] = true
		}
	}
}

func solution(event *Event) (parsedEvent *Event, err error) {
	switch os.Getenv("SOLUTION") {
	case "bloom":
		parsedEvent, err = bloomSolution(event)
	case "naive":
		fallthrough
	default:
		parsedEvent, err = naiveMapSolution(event)
	}
	if !errors.Is(err, DuplicateEventError) {
		db.CreateEvent(ctx, database.CreateEventParams{
			Hash:      event.Hash,
			Block:     int32(event.Block),
			Timestamp: int32(event.Timestamp),
			Amount:    event.Amount,
			From:      event.From,
		})
	}
	return parsedEvent, err
}
