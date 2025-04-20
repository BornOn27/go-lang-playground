package main

import (
	"GoPlayground/goRoutine/02-workerPool/consumer"
	"GoPlayground/goRoutine/02-workerPool/producer"
	"fmt"
	"sync"
	"time"
)

func UnBufferedWorkerPoolExecution(producerDelay time.Duration, consumerDelay time.Duration) {
	fmt.Println("UnbufferedWorkerPoolExecution started...")

	// Create a channel for communication
	ch := make(chan string)
	producerCount := 100
	consumerCount := 5

	// Create and start the producer
	producerInstance := producer.NewProducer(producerCount, producerDelay)
	go producerInstance.Produce(ch)

	// Create and start the consumer
	var wg sync.WaitGroup
	wg.Add(consumerCount)

	consumerInstance := consumer.NewConsumer(consumerCount, consumerDelay)
	consumerInstance.Consume(ch, &wg)

	// Wait to let the consumer finish processing
	wg.Wait()

	fmt.Println("UnbufferedWorkerPoolExecution ended...")
}