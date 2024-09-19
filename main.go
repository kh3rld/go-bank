package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var err error
	state, err = LoadState("transactions.json")
	if err != nil {
		state = State{Balance: 1000}
	}

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . [credit|debit] amount")
		return
	}

	tr := os.Args[1]
	am, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	switch tr {
	case "credit":
		Cred(am)
		fmt.Printf("Kherld, your a/c has been credited with %v, your new balance is %v\n", am, state.Balance)
	case "debit":
		if am < state.Balance {
			fmt.Println("Insufficient funds for debit transaction.")
			return
		}
		Deb(am)
		fmt.Printf("Kherld, your a/c has been debited with %v, your new balance is %v\n", am, state.Balance)
	default:
		fmt.Println("Invalid transaction type. Use 'credit' or 'debit'.")
		return
	}
	if err := SaveState("transactions.json", state); err != nil {
		fmt.Printf("Error saving state: %v\n", err)
		return
	}
	printTransactions()
}

func printTransactions() {
	fmt.Println("Transaction history:")
	for _, t := range state.Transactions {
		fmt.Printf("%s: %d\n", t.Type, t.Amount)
	}
}
