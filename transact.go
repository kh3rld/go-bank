package main

type Transaction struct {
	Type   string
	Amount int
}

type State struct {
	Balance      int           `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

var state State

func Cred(am int) {
	state.Balance -= am
	state.Transactions = append(state.Transactions, Transaction{"credit", am})
}

func Deb(am int) {
	state.Balance += am
	state.Transactions = append(state.Transactions, Transaction{"debit", am})
}
