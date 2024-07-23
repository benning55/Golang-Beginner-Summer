package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	incomeExpenseTracker := NewIncomeExpenseTracker()
	err := LoadData(incomeExpenseTracker)
	if err != nil {
		fmt.Println("No previous data found or error loading data. Starting fresh.")
	}

	go AutoSave(incomeExpenseTracker, 1*time.Minute)

	updateChannel := make(chan string)
	go RealTimeCalculation(incomeExpenseTracker, updateChannel)

	displayWelcomeMessage()

	for {
		displayMenu()

		choice := getUserChoice()

		switch choice {
		case 1:
			incomeExpenseTracker.AddIncome()
		case 2:
			incomeExpenseTracker.AddExpense()
		case 3:
			incomeExpenseTracker.ViewSummary()
		case 4:
			incomeExpenseTracker.ViewTransactions()
		case 5:
			incomeExpenseTracker.ViewExpensesByCategory()
		case 6:
			ExitProgram()
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		select {
		case update := <-updateChannel:
			fmt.Println("\nReal-time update:", update)
		default:
			continue
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
	fmt.Println("5. View Expenses By Category")
	fmt.Println("6. Exit")
}

func getUserChoice() int {
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func ExitProgram() {
	fmt.Println("Exiting the program. Goodbye!")
	os.Exit(0)
}
