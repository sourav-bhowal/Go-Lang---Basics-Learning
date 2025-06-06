// Package auth provides authentication functionalities.
package auth

import "fmt"

// login function simulates a user login process.
// It takes a username and password as parameters and prints them.
func Login(username, password string) {	// Capitalized function name for public access from other packages.
	fmt.Println("Logging in with username:", username)
	fmt.Println("Password:", password)
}

// Session struct represents a user session.
type Session struct {
	Username string // Username of the logged-in user.
	Password string // Password of the logged-in user.
}

// NewSession creates a new session for a user.
func NewSession(username, password string) *Session {
	return &Session{
		Username: username,
		Password: password,
	}
}

func GetSessionDetails(session *Session) string {
	// Returns a string representation of the session details.
	return fmt.Sprintf("Session for user: %s", session.Username)
}