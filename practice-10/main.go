package main

import (
	"fmt"
	"sync"
	"time"
)

/*
You are building an API server where you want to limit the number of requests a client can make in a given time window. For example, allow only 3 requests per second per client.
Write a program in Go that uses goroutines and channels to implement a basic rate limiter. Here's how the rate limiter should behave:
    A client can make requests, but if they exceed the rate limit, their requests should be delayed until the limit resets.
    Use a token bucket algorithm for implementing the rate limiter.
*/

// Function to initialize the rate limiter
func initializeRateLimiter(tokensPerSecond, maxTokens int) chan struct{} {
	// Create a channel for tokens
	// Start a goroutine to refill tokens every second
	tokenBucket := make(chan struct{}, maxTokens)

	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(tokensPerSecond))
		defer ticker.Stop()
		for {

			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
					fmt.Println("Bucket is full lets wait")
				}
			}
		}

	}()
	return tokenBucket
}

// Function to simulate a client making requests
func makeRequest(clientID string, tokenBucket chan struct{}, wg *sync.WaitGroup) {
	// Try to acquire a token from the bucket
	// If successful, print that the request was processed
	// If not, delay until a token becomes available

	defer wg.Done()

	select {
	case <-tokenBucket:
		fmt.Printf("%s made a successful request at %s\n", clientID, time.Now().Format("15:04:05.000"))
	default: // No tokens available; request is rate-limited
		fmt.Printf("%s was rate-limited at %s\n", clientID, time.Now().Format("15:04:05.000"))
	}

}

func main() {
	tokensPerSecond := 3
	maxTokens := 5
	tokenBucket := initializeRateLimiter(tokensPerSecond, maxTokens)

	// Simulate multiple clients making requests

	clients := []string{"Client1", "Client2", "Client3"}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		for _, client := range clients {
			wg.Add(1)
			go makeRequest(client, tokenBucket, &wg)
		}
		time.Sleep(200 * time.Millisecond)
	}

	wg.Wait()
	fmt.Println("Rate limiting simulation complete.")
}
