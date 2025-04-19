package profile

import (
	"time"
)

func FetchUserDetails(userName string, d time.Duration) []string {
	// Simulate fetching user details from a database or an API
	userDetails := []string{"Name:Admin", "Password:pass"}
	time.Sleep(d)
	return userDetails
}

func FetchUserDetailsCh(userName string, d time.Duration, resultProfile chan []string) {
	// Simulate fetching user details from a database or an API
	userDetails := []string{"Name:Admin", "Password:pass"}
	time.Sleep(d)
	resultProfile <- userDetails
}