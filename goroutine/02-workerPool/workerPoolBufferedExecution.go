package main

import (
	"GoPlayground/goRoutine/02-workerPool/consumer"
	"GoPlayground/goRoutine/02-workerPool/producer"
	"fmt"
	"sync"
	"time"
)

func BufferedWorkerPoolExecution(producerDelay time.Duration, consumerDelay time.Duration) {

	fmt.Println("BufferedWorkerPoolExecution started...")
	producerCount := 100
	consumerCount := 5
	// Create a buffered channel with a capacity of 100
	ch := make(chan string, 100)

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

	fmt.Println("BufferedWorkerPoolExecution ended...")
}
