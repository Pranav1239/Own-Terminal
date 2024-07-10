// utils.go
package utils

import (
	"fmt"
	"os"
)

// CheckErr prints an error message and exits the program if the error is not nil
func CheckErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// PrintWelcomeMessage prints a welcome message with the provided username
func PrintWelcomeMessage(username string) {
	fmt.Printf("Welcome, %s!\n", username)
}

// ExampleUtilityFunction is a placeholder for future utility functions
func ExampleUtilityFunction() {
	fmt.Println("This is an example utility function.")
}
