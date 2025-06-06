package main

import (
	// "github.com/sourav-bhowal/my_module/auth"
	"fmt"

	"github.com/sourav-bhowal/my_module/user"
)

// This file is part of the Go programming language's package management system.
// It defines the structure and methods for managing packages in a Go project.

func main() {
	// auth.Login("sourav", "password123")
	// session := auth.NewSession("sourav", "password123")
	// println(auth.GetSessionDetails(session))

	us := user.NewUser("Sourav", "1234567890", "sourav@gmail.com")
	fmt.Println(user.GetUserDetails(us)) // This will print the user details to the console.
}
