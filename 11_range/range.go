package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4, 5}

	// // normal for loop
	// for i := range nums {
	// 	println(nums[i])
	// }

	// // Sum of all numbers
	// sum := 0
	// for index, num := range nums { // for range, the first value is the index, the second is the value
	// 	sum += num
	// 	println("Index:", index, "Value:", num)
	// }
	// println("Sum:", sum)

	for i, c := range "EXPRESS" {
		fmt.Println(i,string(c))
	}

	myMap := map[string]int{"a": 1, "b": 2, "c": 3}

	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

}
