package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"

	ssh "github.com/Pranav1239/Own-Terminal/internal/ssh"
	todo "github.com/Pranav1239/Own-Terminal/internal/todo"
	user "github.com/Pranav1239/Own-Terminal/internal/user"
	utils "github.com/Pranav1239/Own-Terminal/pkg/utils"
)

func main() {
	// Display welcome message
	username := user.GetUsername()
	utils.PrintWelcomeMessage(username)
	// utils.RunNeofetch()

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
			helpCommand()

		case "signinToSSH":
			handleSSHCommand(args)

		case "todos":
			handleTodosCommand(scanner)

		case "exit":
			fmt.Println("Exiting Own-terminal.")
			return

		default:
			fmt.Println("Unknown command. Available commands: signinToSSH, todos, exit")
		}
	}

	if err := scanner.Err(); err != nil {
		utils.CheckErr(err)
	}
}

func helpCommand() {
	color.Cyan.Println("     signinToSSH")
	color.Green.Println("     todos")
	color.Blue.Println("     exit")
}

func handleSSHCommand(args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: signinToSSH <host> <user> <password>")
		return
	}
	host := args[1]
	user := args[2]
	password := args[3]
	err := ssh.SigninToSSH(host, user, password)
	utils.CheckErr(err)
}

func handleTodosCommand(scanner *bufio.Scanner) {
	for {
		fmt.Println("Todos options:")
		fmt.Println("1. add")
		fmt.Println("2. view")
		fmt.Println("3. delete")
		fmt.Println("4. back to main commands")
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}
		todoInput := scanner.Text()

		switch todoInput {
		case "1", "add":
			addTodoCommand(scanner)
		case "2", "view":
			viewTodosCommand()
		case "3", "delete":
			deleteTodoCommand(scanner)
		case "4", "back", "exit":
			return
		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}
	}
}

func addTodoCommand(scanner *bufio.Scanner) {
	fmt.Print("Enter task to add: ")
	if !scanner.Scan() {
		return
	}
	task := scanner.Text()
	err := todo.AddTodo(task)
	utils.CheckErr(err)
	fmt.Println("To-do item added.")
}

func viewTodosCommand() {
	todos, err := todo.ViewTodos()
	utils.CheckErr(err)
	for i, task := range todos {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func deleteTodoCommand(scanner *bufio.Scanner) {
	fmt.Print("Enter task number to delete: ")
	if !scanner.Scan() {
		return
	}
	taskID := scanner.Text()
	err := todo.DeleteTodo(taskID)
	utils.CheckErr(err)
	fmt.Println("To-do item deleted.")
}
