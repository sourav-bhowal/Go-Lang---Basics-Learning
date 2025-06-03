package main

import (
	"fmt"
	"time"
)

// Note: This code demonstrates the use of channels in Go.
// Channels are used to communicate between goroutines and synchronize their execution.

// SEND
// processNum is a function that simulates processing a number from a channel
// func processNum(numChan chan int) {
// 	fmt.Println("Processing numbers...")

// 	// Use a for loop to continuously read from the channel
// 	for num := range numChan {
// 		fmt.Println("Received:", num)
// 		time.Sleep(1 * time.Second) // Simulate some processing time
// 	}
// }

// RECEIVE
// sum is a function that calculates the sum of a variable number of integers
// func sum(result chan int, numbers ...int) {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	result <- total // Send the result back to the channel
// }

// Go-Routine synchronization example
// func task(done chan bool) {

// 	defer func() { // Defer is used so that if any error occurs, the channel is closed properly. Its like clean up code.
// 		done <- true
// 	}()

// 	// Simulate some work
// 	fmt.Println("Task is running...")
// }

func main() {
	// Create a channel to receive the result of the sum
	// result := make(chan int)

	// // Start a goroutine to calculate the sum of numbers
	// go sum(result, 1, 2, 3, 4, 5)

	// // Wait for the result from the channel
	// res := <-result

	// fmt.Println("Sum:", res)

	// Create a channel using 'make' to communicate integers
	// numChan := make(chan int)

	// // Start a goroutine to send numbers to the channel
	// go processNum(numChan)

	// // Send numbers to the channel
	// for i := range 5 {
	// 	numChan <- i // This will block until the goroutine reads from the channel
	// }

	// time.Sleep(2 * time.Second) // Wait for a moment to ensure the goroutine has time to process

	// // Create a channel using 'make'
	// messageChannel := make(chan string)

	// // Sending a string 'ping in the channel
	// messageChannel <- "ping" // blocking

	// // Recieving the msgs out from the channel
	// msg := <-messageChannel

	// fmt.Println(msg)

	// Create a channel to signal when the task is done
	// done := make(chan bool)

	// // Start a goroutine to run the task
	// go task(done)

	// <- done	// The recieve channel will block until the task is done, ensuring that the main function waits for the goroutine to finish.

	// NOTE: Till now we have buffer less channels, which means that the sender will block until the receiver is ready to receive the message. //

	// NOTE: Buffered channels allow you to send messages without blocking the sender until the receiver is ready.
	// The buffer size is specified when creating the channel.
	// They will only block when the buffer is full or empty i.e. when it exceeds the buffer size. Here(100) is the buffer size.
	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done) // Start the email sender goroutine

	// // Start a goroutine to send emails
	// for i := range 10 { // Simulating sending 100 emails
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("All emails sent to the channel. Waiting for the email sender to finish...")
	// close(emailChan) // Close the channel to signal that no more emails will be sent
	// <- done

	// emailChan <- "souravbhowal@gmail.com"
	// emailChan <- "sourav@gmail.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	chan1 := make(chan int)    // Create a channel to communicate integers
	chan2 := make(chan string) // Create another channel to communicate strings

	// Start two goroutines to send values to the channels its in closure to the main function.
	go func() {
		chan1 <- 42 // Send an integer to chan1
	}()

	go func() {
		chan2 <- "Hello, World!" // Send a string to chan2
	}()

	// Receive from chan1 and chan2 together using a select statement
	for range 2 {
		select {
		case chan1Value := <-chan1: // Receive from chan1
			fmt.Println("Received from chan1:", chan1Value)
		case chan2Value := <-chan2: // Receive from chan2
			fmt.Println("Received from chan2:", chan2Value)
		}
		// The select statement allows you to wait for multiple channel operations.
	}

}

// emailSender is a function that simulates sending emails.
// It takes a channel to receive email addresses and a done channel to send a signal when done.
// The <-chan syntax indicates that this channel is for receiving values, while the chan<- syntax indicates that this channel is for sending values.
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { // Defer is used so that if any error occurs, the channel is closed properly. Its like clean up code.
		done <- true
	}()
	for email := range emailChan {
		fmt.Println("Sending email to:", email)

		// Simulate sending email
		time.Sleep(1 * time.Second) // Simulate some delay in sending email
	}
}
