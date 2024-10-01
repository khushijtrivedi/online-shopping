package models

import "errors"

// User represents a user in the system.
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// In-memory data store for users.
var users = make(map[string]User)

// RegisterUser adds a new user to the data store.
func RegisterUser(email, password string) error {
	if _, exists := users[email]; exists {
		return errors.New("user already exists")
	}

	users[email] = User{
		Email:    email,
		Password: password,
	}
	return nil
}

// GetUser retrieves a user by their email address.
func GetUser(email string) (User, error) {
	user, exists := users[email]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}
