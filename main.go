package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

// Add new Transaction
func (bt *BudgetTracker) AddTransaction(amount float64, category, tType string) {
	newTransaction := Transaction{
		ID:       bt.nextID,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     tType,
	}
	bt.transactions = append(bt.transactions, newTransaction)
	bt.nextID++
}

// Creating DisplayTransaction Methods
func (bt BudgetTracker) DisplayTransaction() {
	fmt.Println("ID\tAmount\tCategory\tDate\tType")

	for _, transaction := range bt.transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n",
			transaction.ID, transaction.Amount, transaction.Category,
			transaction.Date.Format("2006-01-02"),
			transaction.Type)
	}
}

// Get total income or expense
func (bt BudgetTracker) CalculateTotal(tType string) float64 {
	var total float64
	for _, transaction := range bt.transactions {
		if transaction.Type == tType {
			total += transaction.Amount
		}
	}
	return total
}

func (bt BudgetTracker) SavetoCSV(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Amount", "Category", "Date", "Type"})

	for _, t := range bt.transactions {
		record := []string{
			strconv.Itoa(t.ID),
			fmt.Sprintf("%.2f", t.Amount),
			t.Category,
			t.Date.Format("2006-01-02"),
			t.Type,
		}
		writer.Write(record)
	}

	fmt.Println("Transaction Saved to", filename)

	return nil
}

func main() {

}
