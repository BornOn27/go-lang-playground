package consumer

import (
	"fmt"
	"sync"
	"time"
)

type Consumer struct {
	count int
	consumerDelay time.Duration
}

func NewConsumer(count int, duration time.Duration) *Consumer {
	return &Consumer{
		count: count,
		consumerDelay: duration,
	}
}

func (c *Consumer) Consume(ch <-chan string, wg *sync.WaitGroup) {

	for i := 1; i <= c.count; i++ {
		go c.startConsumer(i, ch, wg)
	}

}

func (c *Consumer) startConsumer(consumer int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		// Process the message
		time.Sleep(c.consumerDelay) // Simulate processing time
		fmt.Println("[", consumer, "] Consumed :", msg)
	}

}