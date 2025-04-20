package main

import (
	"fmt"
	"time"
)

func main() {
	producerDelay := time.Duration(25) * time.Millisecond
	consumerDelay := time.Duration(200) * time.Millisecond

	BufferedWorkerPoolExecution(producerDelay, consumerDelay)
	fmt.Println("=====================:::::::::::::::::::=====================")

	UnBufferedWorkerPoolExecution(producerDelay, consumerDelay)
}
