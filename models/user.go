package models

import (
	"errors"
	"fmt"
)

// User represents a user in the system.
type User struct {
	ID       string `json:"id"`       // Unique user ID
	Email    string `json:"email"`    // User's email
	Password string `json:"password"` // User's password
}

// In-memory data store for users.
var users = make(map[string]User)
var userIDCounter int // Counter to generate unique user IDs

// RegisterUser adds a new user to the data store.
func RegisterUser(email, password string) (string, error) { // Return user ID
	if _, exists := users[email]; exists {
		return "", errors.New("user already exists")
	}

	// Generate a unique user ID
	userIDCounter++
	userID := fmt.Sprintf("user-%d", userIDCounter)

	users[email] = User{
		ID:       userID,
		Email:    email,
		Password: password,
	}
	return userID, nil
}

// GetUser retrieves a user by their email address.
func GetUser(email string) (User, error) {
	user, exists := users[email]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}
