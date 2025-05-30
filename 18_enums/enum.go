package main

import "fmt"

// There is no built-in enum type in Go, but we can use constants to simulate enums using types and constants.

// Defining an enum-like type for order status
type OrderStatus string

// Constants representing different order statuses
const (
	Pending   OrderStatus = "Pending"
	Shipped   OrderStatus = "Shipped"
	Delivered OrderStatus = "Delivered"
	Cancelled OrderStatus = "Cancelled"
)

// Function to change order status with an enum-like parameter
func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to:", status)
}

// Main function to demonstrate the use of the enum-like type
func main() {
	changeOrderStatus(Delivered)
}
