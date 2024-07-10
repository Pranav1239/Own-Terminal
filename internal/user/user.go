// internal/user/user.go
package user

import (
	"os"
	"os/user"
)

// GetUsername retrieves the current user's username
func GetUsername() string {
	// First, try to get the username from the environment variable
	username := os.Getenv("USER")
	if username != "" {
		return username
	}

	// If the environment variable is not set, use os/user package
	currentUser, err := user.Current()
	if err == nil && currentUser.Username != "" {
		return currentUser.Username
	}

	// If all else fails, return a default username
	return "User"
}
