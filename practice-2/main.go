package main

import (
	"fmt"
	"sync"
	"time"
)

/*Task: Simulating Food Order Processing
You are building a basic system for a restaurant where:
    Customers place orders.
    The system processes orders concurrently (using goroutines).
    A delivery team picks up the orders after processing.

Write a Go program to:
    Accept a list of orders.
    Use goroutines to process each order.
    Once processed, send the processed orders to another goroutine for delivery.

Requirements:
    Input: A list of order IDs (e.g., order1, order2, order3).
    Processing:
        Each order takes 1 second to process (simulate with time.Sleep).
    Delivery:
        Once an order is processed, a separate goroutine prints "Delivered order <order ID>".
*/

func processOrders(orderID string, deliveryChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Processing order:", orderID)
	if orderID == "order1" {
		time.Sleep(5 * time.Second)
	}
	deliveryChan <- orderID
}

func deliverOrders(deliveryChan chan string, doneChan chan bool) {
	for order := range deliveryChan {
		fmt.Println("Delivering order:", order)
	}
	doneChan <- true
}

func main() {
	orders := []string{"order1", "order2", "order3"}
	deliveryChan := make(chan string, len(orders))
	doneChan := make(chan bool)

	var wg sync.WaitGroup

	go deliverOrders(deliveryChan, doneChan)

	for _, order := range orders {
		wg.Add(1)
		go processOrders(order, deliveryChan, &wg)
	}

	// wg.Add(1)

	wg.Wait()
	close(deliveryChan)
	<-doneChan

	fmt.Println("All orders processed and delivered!")
}
