# Own-terminal

**Note: This project is still under development.**

Own-terminal is a custom terminal application built with Go, designed to provide a personalized experience with features like displaying the username, SSH sign-in, and a to-do list management system.

## Goals

- **Display Username**: Show the username at the start with various customizable options.
- **Sign in to SSH**: Allow users to sign in to SSH directly from the terminal.
- **Manage To-dos**: Add, view, and delete to-do items within the terminal.

## Features

1. **Display Username**
   - At the start of the terminal session, the application displays the username with customizable options such as color themes and styles.

2. **Sign in to SSH**
   - Easily sign in to SSH servers using the terminal.
   - Manage multiple SSH connections with saved credentials.

3. **To-do List Management**
   - Add to-do items: Users can add new tasks to their to-do list.
   - View to-do items: Display the list of current to-dos.
   - Delete to-do items: Remove tasks from the list when completed.

## Installation

To install and run Own-terminal, ensure you have [Go](https://golang.org/dl/) installed on your machine.

1. Clone the repository:
    ```sh
    git clone https://github.com/Pranav1239/Go-Terminal.git
    cd Go-terminal
    ```

2. Build the application:
    ```sh
    go build -o own-terminal
    ```

3. Run the application:
    ```sh
    ./own-terminal
    ```

## Usage

1. **Display Username**
   - On launching the terminal, your username will be displayed. You can customize the display options in the settings file (`config.json`).

2. **Sign in to SSH**
   - Use the `signinToSSH` command to connect to an SSH server. Follow the prompts to enter your credentials.

3. **To-do List Management**
   - Use the following commands to manage your to-dos:
     - `addTodo "Your task"`: Add a new to-do item.
     - `viewTodos`: Display the list of current to-dos.
     - `deleteTodo <task-id>`: Delete a to-do item by its ID.

## Configuration

Customize the application settings in the `config.json` file to personalize your terminal experience, including username display options and SSH settings.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or feedback, please reach out to us at [assist@plecred.co.in](mailto:assist@plecred.co.in).

---

Enjoy your customized terminal experience with Own-terminal!
