package user

type User struct {
	Username string // Username of the user.
	Password string // Password of the user.
	Email    string // Email of the user.
}

// NewUser creates a new user with the provided username, password, and email.
func NewUser(username, password, email string) *User {
	// Returns a pointer to a new User instance with the provided details.
	// It returns a pointer to avoid copying the entire struct, which can be more efficient.
	return &User{
		Username: username,
		Password: password,
		Email:    email,
	}
}

// GetUserDetails returns a string representation of the user's details.
func GetUserDetails(user *User) string {
	// Returns a string representation of the user details.
	return "User Details: " + user.Username + ", " + user.Email
}