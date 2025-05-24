package main

import (
	"fmt"
	// "slices"
)

// This program demonstrates basic operations on slices in Go.
// Slices are dynamically-sized, flexible views into the elements of an array.

func main() {
	// Create a slice of integers
	// var numbers []int

	// // Print the initial slice (should be nil)
	// fmt.Println("Initial slice:", numbers)

	// // Initialize the slice with some values
	// numbers = []int{1, 2, 3, 4, 5}

	// // Print the original slice
	// fmt.Println("Original slice:", numbers)

	// // New slice with predefined values
	// num2 := []int{6, 7, 8, 9, 10}

	// // Print the second slice
	// fmt.Println("Second slice:", num2)

	// // Append a new number to the slice
	// numbers = append(numbers, 6)
	// fmt.Println("After appending 6:", numbers)

	// // Remove the second element (index 1) to the third element (index 3)
	// numbers = slices.Delete(numbers, 1, 3)
	// fmt.Println("After removing the second element:", numbers)

	// Create a new slice with a capacity of 10 using make function
	newSlice := make([]int, 0, 10)

	// Check capacity and length of the new slice
	// fmt.Println("Length of new slice:", len(newSlice))
	// fmt.Println("Capacity of new slice:", cap(newSlice))

	// Append multiple elements to the new slice
	newSlice = append(newSlice, 11, 12, 13)
	// fmt.Println("New slice after appending:", newSlice)

	// // Modify the first element
	// newSlice[0] = 100
	// fmt.Println("New slice after modifying first element:", newSlice)


	// // Copy the new slice to another slice
	// copiedSlice := make([]int, len(newSlice))
	// // Use the copy function to copy elements from newSlice to copiedSlice
	// copy(copiedSlice, newSlice)
	// // Print the copied slice
	// fmt.Println("Copied slice:", copiedSlice)

	// Slice operations
	fmt.Println("Original slice:", newSlice)
	fmt.Println("First two elements:", newSlice[:2]) // First two elements
	fmt.Println("Last two elements:", newSlice[len(newSlice)-2:]) // Last two elements
	fmt.Println("Middle element:", newSlice[1:2]) // Middle element (slice from index 1 to 2)
	fmt.Println("Slice from index 1 to 3:", newSlice[1:3]) // Slice from index 1 to 3
}
