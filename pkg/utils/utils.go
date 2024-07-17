// utils.go
package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gookit/color"
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
	fmt.Println(color.FgGreen.Render(fmt.Sprintf("Welcome, %s!", username)))
	fmt.Println("Type 'help' for available commands.")
}

// ExampleUtilityFunction is a placeholder for future utility functions
func ExampleUtilityFunction() {
	fmt.Println("This is an example utility function.")
}

// RunNeofetch runs the neofetch command and captures its output
func RunNeofetch() {
	cmd := exec.Command("neofetch")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running neofetch:", err)
		return
	}
	fmt.Println(strings.TrimSpace(string(output)))
}
