package main

import (
	"fmt"
	"sync"	// "sync" package provides basic synchronization primitives such as mutual exclusion locks and wait groups.
	// "time"
)

// NOTE: This code demonstrates the use of goroutines in Go.
// Goroutines are lightweight threads managed by the Go runtime.
// It creates multiple goroutines that execute a simple task and an anonymous function as multithreaded operations.

// Task simulates a simple task that prints its ID. It takes an integer ID and a pointer to a sync.WaitGroup.
func task(id int, wg *sync.WaitGroup) {
	// Ensure that the wait group is done when the goroutine completes
	// 'defer' is used to ensure that the Done method is called when the function exits, regardless of how it exits.
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Printf("Task %d is starting\n", id)
}

// This function simulates some work by sleeping for a short duration.
func main() {

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup	// sync.WaitGroup is used to wait for a collection of goroutines to finish.

	// Launch multiple goroutines to perform the task
	for i := 1; i <= 10; i++ {
		// Add to the wait group for each goroutine. Its a counter that tracks how many goroutines are running.
		wg.Add(1)
		// Launch a goroutine to perform the task
		go task(i, &wg)
	}

	// Wait for all goroutines to finish. So the main goroutine will block until all tasks are done.
	wg.Wait()

	// sleep to allow goroutines to finish otherwise the program may exit before they complete
	// time.Sleep(2 * time.Second)
}
