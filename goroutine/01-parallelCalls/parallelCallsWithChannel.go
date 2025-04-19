package main

import (
	"GoPlayground/goRoutine/01-parallelCalls/order"
	"GoPlayground/goRoutine/01-parallelCalls/profile"
	"GoPlayground/goRoutine/01-parallelCalls/transactions"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func parallelCallsWithChannel() {

	delayDuration := time.Duration(rand.Int()%5) * time.Second
	startTime := time.Now()

	// Simulate a delay
	userName := "John Doe"
	fmt.Println("Fetching user details, orders, and transactions in parallel...")

	// Start goroutines for each function
	var wg sync.WaitGroup
	// Add the number of goroutines to wait for
	wg.Add(3)

	// Create a channel to receive results
	resultProfile := make(chan []string, 1)
	resultOrder := make(chan []string, 1)
	resultTxn := make(chan []string, 1)

	go func() {
		defer wg.Done()
		profile.FetchUserDetailsCh(userName, delayDuration, resultProfile)
		fmt.Println("[",delayDuration,"] UserService ")
	}()

	delayDuration = time.Duration(rand.Int()%5) * time.Second
	go func() {
		defer wg.Done()
		order.FetchOrdersForUserCh(userName, delayDuration, resultOrder)
		fmt.Println("[",delayDuration,"] OrderService")
	}()

	delayDuration = time.Duration(rand.Int()%5) * time.Second
	go func() {
		defer wg.Done()
		transactions.FetchUserTransactionsCh(userName, delayDuration, resultTxn)
		fmt.Println("[",delayDuration,"] TransactionService")
	}()

	fmt.Println("Waiting for all goroutines to finish...")

	wg.Wait()
	close(resultProfile)
	close(resultOrder)
	close(resultTxn)

	// Collect results from channels
	fmt.Println(<-resultProfile)
	fmt.Println(<-resultOrder)
	fmt.Println(<-resultTxn)

	endTime := time.Now()

	fmt.Println("All Finished. TotalTime taken in second: ", endTime.Sub(startTime).Seconds())
}
