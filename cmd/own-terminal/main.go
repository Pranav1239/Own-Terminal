// cmd/own-terminal/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ssh "github.com/Pranav1239/Own-Terminal/internal/ssh"
	todo "github.com/Pranav1239/Own-Terminal/internal/todo"
	user "github.com/Pranav1239/Own-Terminal/internal/user"
	utils "github.com/Pranav1239/Own-Terminal/pkg/utils"
)

func main() {
	// Display welcome message
	username := user.GetUsername()
	utils.PrintWelcomeMessage(username)

	// Command-line interface
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		args := strings.Split(input, " ")
		command := args[0]

		switch command {

		case "help":
			fmt.Print("signinToSSh")
			fmt.Print("addTodo")
			fmt.Print("viewTodos")
			fmt.Print("deleteTodo")
			fmt.Print("exit")

		case "signinToSSH":
			if len(args) < 3 {
				fmt.Println("Usage: signinToSSH <host> <user>")
				continue
			}
			host := args[1]
			user := args[2]
			err := ssh.SigninToSSH(host, user)
			utils.CheckErr(err)

		case "addTodo":
			if len(args) < 2 {
				fmt.Println("Usage: addTodo <task>")
				continue
			}
			task := strings.Join(args[1:], " ")
			err := todo.AddTodo(task)
			utils.CheckErr(err)
			fmt.Println("To-do item added.")

		case "viewTodos":
			todos, err := todo.ViewTodos()
			utils.CheckErr(err)
			for i, task := range todos {
				fmt.Printf("%d. %s\n", i+1, task)
			}

		case "deleteTodo":
			if len(args) < 2 {
				fmt.Println("Usage: deleteTodo <task-id>")
				continue
			}
			taskID := args[1]
			err := todo.DeleteTodo(taskID)
			utils.CheckErr(err)
			fmt.Println("To-do item deleted.")

		case "exit":
			fmt.Println("Exiting Own-terminal.")
			return

		default:
			fmt.Println("Unknown command. Available commands: signinToSSH, addTodo, viewTodos, deleteTodo, exit")
		}
	}

	if err := scanner.Err(); err != nil {
		utils.CheckErr(err)
	}
}
