package main

import (
	"fmt"
	"time"
)

type TransactionType int

const (
	Income TransactionType = iota
	Expense
)

type Transaction struct {
	Type     TransactionType
	Amount   float64
	Date     time.Time
	Category string
}

type FinancialOperation interface {
	Execute() error
	GetAmount() float64
}

type IncomeTransaction struct {
	Transaction
}

type ExpenseTransaction struct {
	Transaction
}

func (i IncomeTransaction) Execute() error {
	if i.Amount < 0 {
		return fmt.Errorf("income amount cannot be negative")
	}
	return nil
}

func (e ExpenseTransaction) Execute() error {
	if e.Amount < 0 {
		return fmt.Errorf("expense amount cannot be negative")
	}
	return nil
}

func (t Transaction) GetAmount() float64 {
	return t.Amount
}
