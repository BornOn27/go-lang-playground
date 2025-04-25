package main

import (
	"fmt"
	"time"
)

func throttleAPI(apiRequests <-chan string, done <-chan struct{}) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			fmt.Println("Shutting down throttler...")
			return

		case <-ticker.C:
			select {
			case req := <-apiRequests:
				callAPI(req)
			default:
				// No request in queue
			}
		}
	}
}
