package main

import (
	"fmt"
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
		category := getCategory()
		TotalExpenses += amount
		fmt.Printf("Expense of %.2f added successfully.\n", amount)
		Transactions = append(Transactions, 0-amount)
		ExpenseCategories[category] += amount
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

func ViewExpensesByCategory() {
	fmt.Println("\n--- Expenses by Category ---")
	for category, total := range ExpenseCategories {
		fmt.Printf("%s: $%.2f\n", category, total)
	}
}

func AddTransaction(amount float64) {
	Transactions = append(Transactions, amount)
}

func getCategory() string {
	fmt.Println("\n--- Expense Category ---")
	fmt.Println("1. Food")
	fmt.Println("2. Transport")
	fmt.Println("3. Other")
	var category int
	fmt.Print("Enter expense category: ")
	fmt.Scanln(&category)

	switch category {
	case 1:
		return "Food"
	case 2:
		return "Transport"
	case 3:
		return "Others"
	}

	return "Unknown"
}
