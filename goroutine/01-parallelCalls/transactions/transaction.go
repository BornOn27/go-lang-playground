package transactions

import (
	"time"
)

func FetchUserTransactions(userName string, d time.Duration) []string {
	time.Sleep(d)
	transaction := []string{"transaction1", "transaction2", "transaction3"}
	return transaction
}

func FetchUserTransactionsCh(userName string, d time.Duration, result chan []string) {
	time.Sleep(d)
	result <- []string{"transaction1", "transaction2", "transaction3"}
}
