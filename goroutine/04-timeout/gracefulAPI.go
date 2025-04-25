package main

import (
	"context"
	"fmt"
	"time"
)

func simulateApiCallGracefully(ctx context.Context, channel chan bool) {
	go callApi(channel)

	select {
	case <-channel:

	case <-ctx.Done():
		fmt.Println("<<< API call time-out gracefully handled !!! >>>")
		return
	}
}


func callApi(channel chan bool)  {
	fmt.Println("Calling API...")
	time.Sleep(3 * time.Second)
	fmt.Println("API call completed")
	channel <- true
}
