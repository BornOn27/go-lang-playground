package main

import (
	"context"
	"log"
)

type Event struct {
	UserID string
	Type   string
	Raw    string
}

type EnrichedEvent struct {
	Event
	UserName string
}

func simpleEnrichStage(ctx context.Context, in <-chan Event) <-chan EnrichedEvent {
	out := make(chan EnrichedEvent)
	go func() {
		defer close(out)
		for e := range in {
			// Fake enrichment
			username := "User-" + e.UserID
			out <- EnrichedEvent{Event: e, UserName: username}
		}
	}()
	return out
}

func SimplePipeline(ctx context.Context) {

	// Stage 1: Ingest raw data
	raw := make(chan string)
	go func() {
		defer close(raw)
		for _, r := range []string{
			`{"UserID":"1", "Type":"login"}`,
			`{"UserID":"2", "Type":"logout"}`,
			`{"UserID":"3", "Type":"login"}`,
		} {
			raw <- r
		}
	}()

	// Stage 2: Parse JSON
	parsed := parseStage(ctx, raw)

	// Stage 3: Filter only login events
	filtered := filterLoginStage(ctx, parsed)

	// Stage 4: Enrich with user metadata
	enriched := simpleEnrichStage(ctx, filtered)

	// Stage 5: Output
	for e := range enriched {
		log.Printf("OUTPUT: %+v\n", e)
	}
}
