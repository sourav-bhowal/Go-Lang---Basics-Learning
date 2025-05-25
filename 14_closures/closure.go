package main

// Closures are functions that capture the surrounding state
// and can access variables from their enclosing scope even after that scope has finished executing.
import "fmt"

// counter returns a closure that increments and returns a count each time it is called.
func counter() func() int {
	count := 0 // This variable is captured by the closure
	// The closure will have access to this variable even after counter() has returned.
	return func() int {
		count++      // Increment the captured variable
		return count // Return the current count
	}
}

func main() {
	countFunc := counter() // Create a new counter closure

	// Call the closure multiple times
	fmt.Println(countFunc()) // Output: 1
	fmt.Println(countFunc()) // Output: 2
	fmt.Println(countFunc()) // Output: 3

	// Each call to countFunc increments the captured count variable
}
