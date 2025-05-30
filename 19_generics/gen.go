package main

import "fmt"

// Defining a generic function that takes a array of type comparable and prints each element
func printArray[T comparable](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

// Doing generic in struct
type Stack[T int] struct {
	elements []T // Generic slice to hold elements of type T
}

// Main function to demonstrate the use of the generic function
func main() {
	// Using the generic function with an array of integers
	// intArray := []int{1, 2, 3, 4, 5}
	// printArray(intArray)

	// // Using the generic function with an array of strings
	// stringArray := []string{"Hello", "World", "Generics", "in", "Go"}
	// printArray(stringArray)

	// // Using the generic function with an array of floats
	// floatArray := []float64{1.1, 2.2, 3.3}
	// printArray(floatArray)

	// Demonstrating the use of a generic struct
	stack := Stack[int]{}

	// Adding elements to the stack
	stack.elements = append(stack.elements, 100, 200, 300)

	fmt.Println("Stack elements:", stack.elements) // Output: Stack elements: [1 2 3]

	stack2 := Stack[int]{
		elements: []int{1, 2, 3, 4, 5},
	}

	printArray(stack2.elements) // Output: 1 2 3 4 5
}
