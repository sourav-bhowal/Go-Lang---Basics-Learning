package main

// Structs are used to group related data together.
// They are similar to classes in other languages but without methods.

import (
	"fmt"
	"time"
)

// Customer is a struct that represents a customer with a name and age.
type Customer struct {
	Name string // Name of the customer
	Age  int    // Age of the customer
}

// Define a struct named Person requiring fields for name, age, and address.
// The Name starts with a capital letter to make it exported, meaning it can be accessed from other packages.
type person struct {
	Name      string
	Age       int
	Address   string
	createdAt time.Time // Adding a field to store the creation time
	Customer
}

// New constructor function to create a new person instance with customer details.
func NewCustomer(name string, age int) *Customer {
	// This function creates a new Customer instance and returns it.
	return &Customer{
		Name: name, // Initialize the Customer field with the given name
		Age:  age,  // Initialize the Customer field with the given age
	}
}

// Constructors are not explicitly defined in Go like in some other languages.
// Instead, we can create a function that returns a new instance of the struct.
func NewPerson(name string, age int, address string) *person {
	// This function creates a new Person instance and returns it.
	per := person{
		Name:      name,
		Age:       age,
		Address:   address,
		createdAt: time.Now(), // Set the creation time to the current time
		Customer:  *NewCustomer(name, age), // Initialize the Customer field using the constructor function
	}

	return &per // Return a pointer to the new person instance
}



// Methods are not defined in the struct itself, but we can define functions that operate on the struct.
// In Go, methods can be defined on types, including structs, but for simplicity, we will just use functions here.

// func (*reciever type) methodName(parameters) returnType {}
// The receiver type is the type of the struct we are defining the method for.
// The person should be a pointer receiver if we want to modify the struct's fields.
func (p *person) changeAge(age int) {
	// This function changes the age of the person.
	p.Age = age // Update the age field of the person struct. We dont need derefernce because structs takes care of that
}

// NOTE: use * is compulsory if we want to modify the struct's fields.
// If we don't use a pointer receiver, we will be working with a copy of the struct, and changes won't affect the original struct.
// NOTE: If we want to read the struct's fields without modifying them, we can use a value receiver (without *).

func (p person) getAge() int {
	// This function returns the age of the person.
	return p.Age // Return the age field of the person struct.
}

func main() {
	// Create a new Person instance
	p1 := person{
		Name:    "Alice",
		Age:     30,
		Address: "123 Main St",
		Customer: Customer{
			Name: "June", // Initialize the Customer field with the same name
			Age:  20,      // Initialize the Customer field with the same age
		},
	}

	// Print the person's details
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)
	fmt.Println("Address:", p1.Address)
	fmt.Println("Customer Name:", p1.Customer.Name)
	fmt.Println("Customer Age:", p1.Customer.Age)

	// p1.changeAge(35) // Change the age using the method

	// // Print the person's details
	// fmt.Println("Name:", p1.Name)
	// fmt.Println("Age from method:", p1.getAge()) // Get the age using the method

	// // Print the person's details
	// fmt.Println("Name:", p.Name)
	// fmt.Println("Age:", p.Age)
	// fmt.Println("Address:", p.Address)
	// fmt.Println("Created At:", p.createdAt)

	// // Modify the person's age
	// p.Age = 31
	// fmt.Println("Updated Age:", p.Age)

	// p.createdAt = time.Now() // Set the creation time to the current time
	// fmt.Println("Updated Created At:", p.createdAt)

	// person2 := person{
	// 	Name:    "Bob",
	// 	Age:     25,
	// 	Address: "456 Elm St",
	// 	createdAt: time.Now(), // Set the creation time to the current time
	// }
	// // Print the details of person2
	// fmt.Println("Name:", person2.Name)
	// fmt.Println("Age:", person2.Age)
	// fmt.Println("Address:", person2.Address)
	// fmt.Println("Created At:", person2.createdAt)

	// p := NewPerson("Alice", 30, "123 Main St") // Create a new person using the constructor function

	// c:= NewCustomer("June", 20) // Create a new customer using the constructor function
	// // Print the person's details
	// fmt.Println("Name:", p.Name)
	// fmt.Println("Age:", p.getAge()) // Get the age using the method


	// fmt.Println("Name:", p.Name)
	// fmt.Println("Age:", p.getAge()) // Get the age using the method
	// fmt.Println("Address:", p.Address)

	// p.changeAge(35)                         // Change the age using the method
	// fmt.Println("Updated Age:", p.getAge()) // Get the updated age using the method

	// This is an example of a struct with fields that are not exported. Its similar to private fields in other languages.
	// language := struct {
	// 	Name    string
	// 	Version string
	// }{
	// 	Name:    "Go",
	// 	Version: "1.20",
	// }

	// fmt.Println("Language Name:", language.Name)
	// fmt.Println("Language Version:", language.Version)
}
