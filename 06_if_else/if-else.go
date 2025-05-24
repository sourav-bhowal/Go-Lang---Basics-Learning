package main

import "fmt"

func main() {
	// User input
	var input int
	fmt.Println("Enter your age: ")
	fmt.Scan(&input)

	if input > 18 {	// if input is greater than 18
		fmt.Println("You are an adult")	
	} else if input < 18 {	// if input is less than 18
		fmt.Println("You are a minor")
	} else {	// if input is equal to 18
		fmt.Println("You are a teenager")
	}

	// Short hand if else 
	if age := 18; age >= 18 {
		fmt.Println("You are an adult", age)
	} else {
		fmt.Println("You are a minor", age)
	}
}