// internal/user/user.go
package user

import (
	"os"
	"os/user"
)

func GetUsername() string {
	username := os.Getenv("USER")
	if username != "" {
		return username
	}

	currentUser, err := user.Current()
	if err == nil && currentUser.Username != "" {
		return currentUser.Username
	}

	return "User"
}
