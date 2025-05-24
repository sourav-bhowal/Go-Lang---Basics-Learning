package main

import "fmt"

// Arrays are fixed-size collections of elements of the same type.
// They are useful when you know the number of elements in advance.
// Arrays are value types, meaning that when you assign an array to another array, a copy of the original array is made.

func main() {
	// Array of fixed size
	var a [5]int
	a[4] = 100     // Assigning value to the 5th element of the array
	fmt.Println(a) // Output: [0 0 0 0 100] all other elements are 0 because they are not initialized

	// Array of strings
	var b [3]string
	b[0] = "Hello"
	b[1] = "World"

	fmt.Println(b) // Output: [Hello World ]  // The last element is empty because we didn't assign a value to it

	// Array with values
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[0])   // Accessing the first element of the array
	fmt.Println(len(nums)) // Length of the array

	// 2D Array
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix) // Output: [[1 2 3] [4 5 6] [7 8 9]]
}

