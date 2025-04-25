package main

import "fmt"

func simulateFaultyApiCall(channel chan bool) {
	fmt.Println("Simulating API call without timeout handling")
	callApi(channel)
}
