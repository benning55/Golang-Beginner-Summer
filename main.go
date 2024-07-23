package main

import "fmt"

var TotalIncome float64
var TotalExpenses float64
var Transactions []float64

func main() {
	displayWelcomeMessage()

	for {
		displayMenu()

		choice := getUserChoice()

		switch choice {
		case 1:
			AddIncome()
		case 2:
			AddExpense()
		case 3:
			ViewSummary()
		case 4:
			ViewTransactions()
		case 5:
			ExitProgram()
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func displayWelcomeMessage() {
	fmt.Println("Welcome to the Income Expenses Tracker!")
	fmt.Println("This program will help you manage your finances.")
}

func displayMenu() {
	fmt.Println("\n--- Income Expense Tracker ---")
	fmt.Println("1. Add Income")
	fmt.Println("2. Add Expense")
	fmt.Println("3. View Summary")
	fmt.Println("4. View Transactions")
	fmt.Println("5. Exit")
}

func getUserChoice() int {
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}
