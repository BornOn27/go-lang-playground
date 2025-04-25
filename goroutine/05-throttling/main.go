package main

import (
	"fmt"
	"time"
)

func main() {
	apiRequests := make(chan string)
	done := make(chan struct{})

	// Start the throttler
	go throttleAPI(apiRequests, done)

	// Simulate rapid incoming API calls
	for i := 1; i <= 5; i++ {
		apiRequests <- fmt.Sprintf("Call-%d", i)
		fmt.Println("Queued request:", i)
	}

	// Give it some time to process before exiting
	time.Sleep(6 * time.Second)
	close(done)
}


func callAPI(req string) {
	// Simulate the actual API call
	fmt.Println("✅ API request sent:", req)
}