package order

import (
	"time"
)

func FetchOrdersForUser(userName string, d time.Duration) []string {
	time.Sleep(d)

	// Simulate fetching orders from a database or an API
	return []string{"Order1", "Order2", "Order3"}
}

func FetchOrdersForUserCh(userName string, d time.Duration, result chan []string) {
	time.Sleep(d)

	// Simulate fetching orders from a database or an API
	result <- []string{"Order1", "Order2", "Order3"}
}
