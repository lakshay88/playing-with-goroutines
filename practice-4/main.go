package main

import (
	"fmt"
	"time"
)

/*
You need to create a worker pool system in Go to process a set of tasks concurrently. Each worker should pick a task, process it, and return the result. Ensure the main function waits until all tasks are completed.
*/

const numWorkers = 3 // Number of workers
const numTasks = 20  // Number of tasks

func worker(workerId int, taskChan chan int, doneChan chan bool) {
	for task := range taskChan {
		fmt.Printf("Worker %d is processing task %d\n", workerId, task)
		time.Sleep(500 * time.Millisecond)
	}

	doneChan <- true
}

func main() {
	taskChan := make(chan int, numTasks)

	doneChan := make(chan bool)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, taskChan, doneChan)
	}

	for task := 1; task <= numTasks; task++ {
		taskChan <- task
	}

	close(taskChan)

	for i := 0; i < numWorkers; i++ {
		<-doneChan
	}

	fmt.Println("All tasks completed!")
}
