package producer

import (
	"fmt"
	"time"
)

type Producer struct {
	count int
	producerDelay time.Duration
}

func NewProducer(count int, producerDelay time.Duration) *Producer {
	return &Producer{
		count: count,
		producerDelay: producerDelay,
	}
}


func (p *Producer) Produce(ch chan<- string) {
	fmt.Println("Starting Producer...")

	for i := 0; i < p.count; i++ {
		ch <- "Message " + fmt.Sprintf("%d", i+1) // Send message to channel
		time.Sleep(p.producerDelay) // Simulate processing time
	}
	close(ch) // Close the channel to signal completion
	fmt.Println("Producer finished sending messages.")
}
