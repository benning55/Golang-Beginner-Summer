# Student Work for Week 1: Introduction to Go & Setting Up Environment

Welcome to the first week of the Go Programming Course! This document outlines the tasks you need to complete as part of your student work for this week.

## Tasks

### 1. Install Go and Set Up Your Development Environment
- **Download Go:**
  - Visit the official Go website: [golang.org](https://golang.org/dl/)
  - Choose the appropriate installer for your operating system (Windows, macOS, Linux) and follow the installation instructions.

- **Verify Installation:**
  - Open your terminal or command prompt.
  - Run the following command to check if Go is installed correctly:
    ```bash
    go version
    ```
  - You should see the installed Go version printed in the terminal.

- **Set Up Your IDE:**
  - **Recommended IDE:** Visual Studio Code
    - Download from [code.visualstudio.com](https://code.visualstudio.com/)
    - Install the Go extension for Visual Studio Code:
      - Open VS Code
      - Go to Extensions (Ctrl+Shift+X)
      - Search for "Go" and install the extension by the Go team
  - **Alternative Options:** GoLand, Sublime Text, Atom, or any text editor of your choice.

### 2. Create and Set Up the Project
- **Create a New Project Directory:**
  - Open your terminal or command prompt.
  - Run the following commands to create and navigate to your project directory:
    ```bash
    mkdir income-expenses-tracker
    cd income-expenses-tracker
    ```

- **Initialize a Go Module:**
  - Run the following command to initialize a new Go module:
    ```bash
    go mod init github.com/yourusername/income-expenses-tracker
    ```
  - Replace `yourusername` with your actual GitHub username or any identifier you prefer.

- **Create the `main.go` File:**
  - In your project directory, create a new file named `main.go`.

### 3. Write Your First Go Program
- **Open `main.go` in Your IDE/Text Editor:**
  - Add the following code to `main.go`:
    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Welcome to the Income Expenses Tracker!")
        fmt.Println("This program will help you manage your finances.")
    }
    ```

### 4. Run Your Program
- **Use the Terminal to Run Your Program:**
  - In your terminal, make sure you are in the project directory.
  - Run the program using the following command:
    ```bash
    go run main.go
    ```
  - You should see the output:
    ```
    Welcome to the Income Expenses Tracker!
    This program will help you manage your finances.
    ```

### 5. Build GO Project
  - Try compiling your program with:
    ```bash
    go build
    ```
  - This will create an executable file in your project directory.
  - Run the executable (on Windows, use `income-expenses-tracker.exe`, on macOS/Linux use `./income-expenses-tracker`).

## Additional Resources
- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)


----