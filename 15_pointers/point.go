package main

import "fmt"

// Point represents a point in 2D space.
// It is defined as a struct with two fields: X and Y.

// func changeNum(num int) {
// 	num = 10
// 	fmt.Println("Before changeNum:", num)
// }

// Pass by reference allows us to modify the original variable.
func changeNum(num *int) {
	*num = 5 // Dereference the pointer to change the value it points to.
	fmt.Println("Before changeNum:", *num)
}

func main() {
	// num:= 5
	// changeNum(num)
	// fmt.Println("After changeNum:", num)

	num := 1
	changeNum(&num)                      // Pass the address of num to changeNum
	fmt.Println("After changeNum:", num) // This will print 5 because we modified the original variable through the pointer.
}
