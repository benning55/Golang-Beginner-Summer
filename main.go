package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Income Expenses Tracker!")
	fmt.Println("This program will help you manage your finances.")

	var totalIncome float64
	var totalExpenses float64

	for {
		fmt.Println("\n--- Income Expense Tracker ---")
		fmt.Println("1. Add Income")
		fmt.Println("2. Add Expense")
		fmt.Println("3. View Summary")
		fmt.Println("4. Add Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter income amount: ")
			var amount float64
			fmt.Scanln(&amount)
			if amount >= 0 {
				totalIncome += amount
				fmt.Printf("Income of %.2f added successfully.\n", amount)
			} else {
				fmt.Println("Invalid amount. Please enter positive income")
			}
		case 2:
			fmt.Println("Enter expense amount: ")
			var amount float64
			fmt.Scanln(&amount)
			if amount >= 0 {
				totalExpenses += amount
				fmt.Printf("Expense of %.2f added successfully.\n", amount)
			} else {
				fmt.Println("Invalid amount. Please enter positive expense")
			}
		case 3:
			fmt.Println("\n--- Financial Summary ---")
			fmt.Printf("Total Income: $%.2f\n", totalIncome)
			fmt.Printf("Total Expenses: $%.2f\n", totalExpenses)
			fmt.Printf("Net Savings: $%.2f\n", totalIncome-totalExpenses)
		case 4:
			fmt.Println("Exiting the program. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
