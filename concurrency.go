package main

import (
	"fmt"
	"time"
)

func AutoSave(tracker *IncomeExpenseTracker, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		err := SaveData(tracker)
		if err != nil {
			fmt.Printf("Auto-save error: %v\n", err)
		} else {
			fmt.Println("Auto-saved successfully")
		}
	}
}

func RealTimeCalculation(tracker *IncomeExpenseTracker, updateChan chan<- string) {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		totalIncome := tracker.calculateTotalIncome()
		totalExpenses := tracker.calculateTotalExpenses()
		netSavings := totalIncome - totalExpenses
		updateChan <- fmt.Sprintf("Income: $%.2f, Expenses: $%.2f, Net Savings: $%.2f",
			totalIncome, totalExpenses, netSavings)
	}
}
