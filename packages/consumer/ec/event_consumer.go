package ec

import (
	"log"
	"time"

	"github.com/alehkonan/sgb/packages/processors"
)

type Consumer struct {
	fetcher   processors.Fetcher
	handler   processors.Handler
	batchSize int
}

// Creates a new consumer that works with events
func New(
	fetcher processors.Fetcher,
	handler processors.Handler,
	batchSize int,
) Consumer {
	return Consumer{fetcher, handler, batchSize}
}

// Starts the event consumer
func (c *Consumer) Start() error {
	for {
		events, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s: ", err.Error())
			continue
		}

		if len(events) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(events); err != nil {
			log.Print(err)
		}
		continue
	}
}

func (c *Consumer) handleEvents(events []processors.Event) error {
	// TODO add sync.WaitGroup
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.handler.Handle(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())
			continue
		}
	}

	return nil
}
