package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	SimplePipeline(ctx)
	fmt.Println("=====================:::::::::::::::::::=====================")
	FanInFanOut(ctx)
}
