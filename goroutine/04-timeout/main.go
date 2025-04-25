package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 2 * time.Second)
	channel := make(chan bool)


	// Simulate API call with timeout
	simulateApiCallGracefully(ctx, channel)

	//fmt.Println("=====================:::::::::::::::::::=====================")

	// Simulate API call without timeout
	//go simulateFaultyApiCall(channel)
	select {
		case <- ctx.Done():
			fmt.Println("<<< Timeout >>> Main Exit !!!")
	}

}
