package main

import (
	"fmt"
	"maps"
)

// Map is a collection of key-value pairs.
// It is implemented as a wrapper around the built-in map type.

// `Make` is a built-in function that initializes a map.

func main() {
	// Creating a map using the built-in make function
	myMap := make(map[string]int)	// make(map[keyType]valueType)

	// Adding key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 3
	myMap["orange"] = 8

	fmt.Println("Initial map:", myMap)

	// Creating a map with initial values using a map literal without make
	myMap2 := map[string]int{
		"grape":  4,
		"kiwi":   6,
		"mango":  2,
		"papaya": 7,
	}
	fmt.Println("Updated map:", myMap2)

	// Accessing values by key
	value, exists := myMap["banana"]	// value, bool:= myMap[key]

	if exists {
		fmt.Printf("Value for 'banana': %d\n", value)
	} else {
		fmt.Println("'banana' not found in the map")
	}

	// deleting a key-value pair
	delete(myMap, "orange")
	fmt.Println("Map after deleting 'orange':", myMap)

	// clearing the map
	clear(myMap2)
	fmt.Println("Map after clearing:", myMap2)

	// Comparing two maps
	fmt.Println("Comparing maps:", maps.Equal(myMap, myMap2))
}
