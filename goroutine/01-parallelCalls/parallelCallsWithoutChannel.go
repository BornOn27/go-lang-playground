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

func parallelCallsWithOutChannel() {

	delayDuration := time.Duration(rand.Int()%5) * time.Second
	startTime := time.Now()

	// Simulate a delay
	userName := "John Doe"
	fmt.Println("Fetching user details, orders, and transactions in parallel...")

	// Start goroutines for each function
	var wg sync.WaitGroup
	// Add the number of goroutines to wait for
	wg.Add(3)


	go func() {
		defer wg.Done()
		userDetails := profile.FetchUserDetails(userName, delayDuration)
		fmt.Println("[",delayDuration,"] UserService :: User details for user", userName, ":", userDetails)
	}()

	delayDuration = time.Duration(rand.Int()%5) * time.Second
	go func() {
		defer wg.Done()
		orders := order.FetchOrdersForUser(userName, delayDuration)
		fmt.Println("[",delayDuration,"] OrderService :: Orders for user", userName, ":", orders)
	}()

	delayDuration = time.Duration(rand.Int()%5) * time.Second
	go func() {
		defer wg.Done()
		transaction := transactions.FetchUserTransactions(userName, delayDuration)
		fmt.Println("[",delayDuration,"] TransactionService :: User transactions for user", userName, ":", transaction)
	}()

	wg.Wait()
	fmt.Println("Waiting for all goroutines to finish...")

	endTime := time.Now()

	fmt.Println("All Finished. TotalTime taken in second: ", endTime.Sub(startTime).Seconds())
}
