package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func fanInFanOutEnrichStage(ctx context.Context, in <-chan Event, workerCount int) <-chan EnrichedEvent {
	out := make(chan EnrichedEvent)
	workerInput := make(chan Event)

	var wg sync.WaitGroup

	// Start N workers (fan-out)
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go enrichWorker(ctx, i, workerInput, out, &wg)
	}

	// Feed the input channel to workers
	go func() {
		for e := range in {
			workerInput <- e
		}
		close(workerInput) // close so workers know to stop
	}()

	// Close output after all workers are done (fan-in)
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func enrichWorker(ctx context.Context, id int, in <-chan Event, out chan<- EnrichedEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	for e := range in {
		// Simulate slow API call
		time.Sleep(200 * time.Millisecond)

		username := "User-" + e.UserID
		out <- EnrichedEvent{Event: e, UserName: username}
	}
}


func FanInFanOut(ctx context.Context) {
	raw := make(chan string)

	// Simulate input stream
	go func() {
		defer close(raw)
		for _, r := range []string{
			`{"UserID":"1", "Type":"login"}`,
			`{"UserID":"2", "Type":"logout"}`,
			`{"UserID":"3", "Type":"login"}`,
			`{"UserID":"4", "Type":"login"}`,
		} {
			raw <- r
		}
	}()

	parsed := parseStage(ctx, raw)
	filtered := filterLoginStage(ctx, parsed)

	// Fan-out 3 workers
	enriched := fanInFanOutEnrichStage(ctx, filtered, 3)

	for e := range enriched {
		log.Printf("OUTPUT: %+v\n", e)
	}
}

