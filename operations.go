package main

import (
	"fmt"
	"time"
)

type IncomeExpenseTracker struct {
	TotalIncome       float64
	TotalExpenses     float64
	Transactions      []Transaction
	ExpenseCategories map[string]float64
}

func NewIncomeExpenseTracker() *IncomeExpenseTracker {
	return &IncomeExpenseTracker{
		Transactions:      []Transaction{},
		ExpenseCategories: make(map[string]float64),
	}
}

func (t *IncomeExpenseTracker) AddIncome() {
	amount := getAmount("income")
	income := IncomeTransaction{
		Transaction: Transaction{
			Type:   Income,
			Amount: amount,
			Date:   time.Now(),
		},
	}
	err := income.Execute()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	t.TotalIncome += amount
	t.Transactions = append(t.Transactions, income.Transaction)
	fmt.Printf("Income of $%.2f added successfully.\n", amount)
}

func (t *IncomeExpenseTracker) AddExpense() {
	amount := getAmount("expense")
	category := getCategory()
	expense := ExpenseTransaction{
		Transaction: Transaction{
			Type:     Expense,
			Amount:   amount,
			Date:     time.Now(),
			Category: category,
		},
	}
	err := expense.Execute()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	t.TotalExpenses += amount
	t.Transactions = append(t.Transactions, expense.Transaction)
	t.ExpenseCategories[category] += amount
	fmt.Printf("Expense of $%.2f in category '%s' added successfully.\n", amount, category)
}

func (t *IncomeExpenseTracker) ViewSummary() {
	fmt.Println("\n--- Financial Summary ---")
	fmt.Printf("Total Income: $%.2f\n", t.TotalIncome)
	fmt.Printf("Total Expenses: $%.2f\n", t.TotalExpenses)
	fmt.Printf("Net Savings: $%.2f\n", t.TotalIncome-t.TotalExpenses)
}

func (t *IncomeExpenseTracker) ViewTransactions() {
	fmt.Println("\n--- Financial Transactions ---")
	for index, tx := range t.Transactions {
		if tx.Type == Income {
			fmt.Printf("%d. Income: $%.2f\n", index+1, tx.GetAmount())
		} else {
			fmt.Printf("%d. Expense: $%.2f\n", index+1, tx.GetAmount())
		}
	}
	fmt.Printf("Net Balance: $%.2f\n", t.TotalIncome-t.TotalExpenses)
}

func (t *IncomeExpenseTracker) ViewExpensesByCategory() {
	fmt.Println("\n--- Expenses by Category ---")
	for category, total := range t.ExpenseCategories {
		fmt.Printf("%s: $%.2f\n", category, total)
	}
}

func getAmount(transactionType string) float64 {
	var amount float64
	fmt.Printf("Enter %s amount: ", transactionType)
	fmt.Scanln(&amount)
	return amount
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
