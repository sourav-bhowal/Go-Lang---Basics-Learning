package main

// sum is a variadic function that takes a variable number of integer arguments 
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// any type variadic function
func sumAny(nums ...any) int {
	// Total will hold the sum of integers
	total := 0
	// Iterate over the variadic arguments
	for _, num := range nums {
		// Check if the current argument is of type int
		if n, ok := num.(int); ok {
			// If it is, add it to the total
			total += n
		}
	}
	// If the argument is not an int, it will be ignored
	// Return the total sum of integers
	return total
}

func main() {
	// Example usage
	result := sum(1, 2, 3, 4, 5)
	println("The sum is:", result) // Output: The sum is: 15

	resultAny := sumAny(1, 2, "three", 4.0, 5)

	println("The sum of integers is:", resultAny) // Output: The sum of integers is: 12
}