package main

import "fmt"

// Task: Sum of Numbers Using Multiple Goroutines

// Write a Go program to calculate the sum of an array of numbers by dividing the array into chunks and processing each chunk in a separate goroutine. Use a channel to collect the partial sums from each goroutine and calculate the final sum in the main function.

func calculateChunkSum(numbers []int, results chan int) {
	sum := 0
	for _, element := range numbers {
		sum = sum + element
	}
	results <- sum
}

func main() {

	numbers := []int{1, 2, 4, 5, 6, 7, 88, 7, 1, 2, 4, 5, 6, 7, 8, 9, 10, 4, 5, 6}
	numGoroutines := 3
	chunkSize := (len(numbers) + numGoroutines - 1) / numGoroutines

	results := make(chan int, len(numbers))

	for i := 0; i < len(numbers); i += chunkSize {
		end := i + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		go calculateChunkSum(numbers[i:end], results)
	}

	finalSum := 0
	for i := 0; i < numGoroutines; i++ {
		finalSum += <-results
	}

	fmt.Println("================================================")
	fmt.Println(finalSum)
	fmt.Println("================================================")

}
