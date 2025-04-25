package main

import (
	"context"
	"encoding/json"
)

func parseStage(ctx context.Context, in <-chan string) <-chan Event {
	out := make(chan Event)
	go func() {
		defer close(out)
		for r := range in {
			var e Event
			json.Unmarshal([]byte(r), &e)
			e.Raw = r
			out <- e
		}
	}()
	return out
}

func filterLoginStage(ctx context.Context, in <-chan Event) <-chan Event {
	out := make(chan Event)
	go func() {
		defer close(out)
		for e := range in {
			if e.Type == "login" {
				out <- e
			}
		}
	}()
	return out
}
