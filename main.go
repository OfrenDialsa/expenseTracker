package main

import (
	"time"
)

// Transaction Struct for each Transaction's data
type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string // "income" or "expense"
}

// BudgetTracker Struct to manage the budget and transactions
type BudgetTracker struct {
	transactions []Transaction
	nextID       int
}

// Interface for the Common Behavior
type FinancialRecord interface {
	GetAmount() float64
	GetType() string
}

// Implement Interface for Transaction
func (t Transaction) GetAmount() float64 {
	return t.Amount
}

func (t Transaction) GetType() string {
	return t.Type
}

func main() {

}
