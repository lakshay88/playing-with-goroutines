package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Create a system in Go where producers generate data and send it to a shared channel, and consumers process the data. Ensure proper synchronization and graceful termination.
Requirements
    Producers:
        There should be P producers, each generating N items.
        Producers should send their generated items to a shared channel.
    Consumers:
        There should be C consumers, each reading from the shared channel and processing data.
        Each consumer should print the data it processes.
    Graceful Termination:
        Use a way to signal when all producers have finished producing.
        Ensure consumers stop processing once all data is consumed.
*/

func producers(producerId, itemsPerProducer int, dataChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= itemsPerProducer; i++ {
		fmt.Printf("Producer %d: produced item %d\n", producerId, i)
		dataChan <- i
		time.Sleep(100 * time.Millisecond)
	}
}

func consumers(id int, dataChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range dataChan {
		time.Sleep(550 * time.Millisecond)
		fmt.Printf("Consumer %d: processed item %d\n", id, item)
	}
}

func main() {
	// Define the number of producers, consumers, and items per producer
	numProducers := 2
	numConsumers := 3
	itemsPerProducer := 5

	// Create a shared channel for data
	dataChan := make(chan int, 10)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	producerWg.Add(numProducers)
	for i := 1; i <= numProducers; i++ {
		go producers(i, itemsPerProducer, dataChan, &producerWg)
	}

	consumerWg.Add(numConsumers)
	for i := 1; i <= numConsumers; i++ {
		go consumers(i, dataChan, &consumerWg)
	}

	go func() {
		producerWg.Wait() // Wait for all producers to finish
		close(dataChan)   // Close the channel
	}()

	consumerWg.Wait()

	fmt.Println("All data processed!")
}
