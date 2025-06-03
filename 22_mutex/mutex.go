package main

import (
	"fmt"
	"sync"
)

// NOTE: Mutexes are used to protect shared data from concurrent access.
// It ensures that only one goroutine can access the shared data at a time.
// Mutex is a mutual exclusion lock that can be used to synchronize access to shared data.

// Type post represents a blog post with a view count.
type post struct {
	views int
	// mutex is used to protect the views field from concurrent access. So to avoid race conditions while incrementing the views.
	mutex sync.Mutex
}

// increment is a method that increments the view count of the post.
func (p *post) increment(wg *sync.WaitGroup) {
	// Defer the call to Done and Unlock to ensure they are called when the function returns.
	defer func() {
		wg.Done()        // Decrement the wait group counter when the function returns.
		p.mutex.Unlock() // Ensure the mutex is unlocked when the function returns, even if it panics.
	}()

	// Lock the mutex to ensure that only one goroutine can access the views field at a time.
	p.mutex.Lock() // This bottlenecks the access to the views field, ensuring that only one goroutine can increment it at a time.

	// Increment the views field of the post.
	p.views += 1
}

// main function is the entry point of the program.
func main() {
	// wait group is used to wait for a collection of goroutines to finish. Its compulsory to use it when we are using goroutines.
	var wg sync.WaitGroup

	// Create a new post with 0 views. If we dont write 0 it will be 0 by default.
	myPost := post{views: 0}

	for range 100 {
		// Add 1 to the wait group counter for each goroutine.
		wg.Add(1)
		// Increment the view count of the post.
		go myPost.increment(
			&wg, // Pass the wait group to the increment method.
		)
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	// Print the view count of the post.
	fmt.Println("Post views:", myPost.views) // Output: Post views: 1
}
