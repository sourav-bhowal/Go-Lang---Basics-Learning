package main

import (
	"fmt"
	"os"
)

// Note: This file is part of a Go module and should be used in conjunction with other files in the module.

func main() {
	// Open a file named "hello.txt" in the current directory
	file, error := os.Open("hello.txt")

	// Check for errors when opening the file
	if error != nil {
		// Panic if there is an error
		panic(error)
	}

	// Successfully opened the file, now we can defer its closure
	fileInfo, error := file.Stat()

	// Check for errors when getting file information
	if error != nil {
		// Panic if there is an error
		panic(error)
	}

	// Print the file information
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("File Size:", fileInfo.Size(), "bytes")
	fmt.Println("File Mode:", fileInfo.Mode())
	fmt.Println("File Modified:", fileInfo.ModTime())

	defer file.Close() // Ensure the file is closed when we're done
	println("File opened successfully, and will be closed now.")
}
