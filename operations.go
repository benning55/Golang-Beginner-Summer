package main

import (
	"fmt"
	"os"
)

func AddIncome() {
	fmt.Print("Enter income amount: ")
	var amount float64
	fmt.Scan(&amount)
	if amount >= 0 {
		TotalIncome += amount
		fmt.Printf("Income of %.2f added successfully.\n", amount)
		Transactions = append(Transactions, amount)
	} else {
		fmt.Println("Invalid amount. Please enter positive income")
	}
}

func AddExpense() {
	fmt.Print("Enter expense amount: ")
	var amount float64
	fmt.Scan(&amount)
	if amount >= 0 {
		TotalExpenses += amount
		fmt.Printf("Expense of %.2f added successfully.\n", amount)
		Transactions = append(Transactions, 0-amount)
	} else {
		fmt.Println("Invalid amount. Please enter positive expense")
	}
}

func ViewSummary() {
	fmt.Println("\n--- Financial Summary ---")
	fmt.Printf("Total Income: $%.2f\n", TotalIncome)
	fmt.Printf("Total Expenses: $%.2f\n", TotalExpenses)
	fmt.Printf("Net Savings: $%.2f\n", TotalIncome-TotalExpenses)
}

func ViewTransactions() {
	fmt.Println("\n--- Financial Transactions ---")
	for index, amount := range Transactions {
		if amount >= 0 {
			fmt.Printf("%d. Income: $%.2f\n", index+1, amount)
		} else {
			fmt.Printf("%d. Expense: $%.2f\n", index+1, amount)
		}
	}
	fmt.Printf("Net Balance: $%.2f\n", TotalIncome-TotalExpenses)
}

func ExitProgram() {
	fmt.Println("Exiting the program. Goodbye!")
	os.Exit(0)
}

func AddTransaction(amount float64) {
	Transactions = append(Transactions, amount)
}
